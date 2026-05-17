package usecases

import (
	"context"
	"fmt"
	"strconv"

	"github.com/twodigitss/reserv-go/internal/modules/payment"
	"github.com/twodigitss/reserv-go/internal/modules/reservation"
	"github.com/twodigitss/reserv-go/internal/modules/tables"
	"github.com/twodigitss/reserv-go/internal/modules/users"
)

type TableReservation struct {
    tables  *tables.TableService
    users   *users.UserService
    payment payment.Provider
    paydb payment.Payment
    reserv reservation.ReservService
}

func NewCreateReservation(
    t *tables.TableService,
    u *users.UserService,
    p payment.Provider,         // ← recibe la interfaz
) *TableReservation {
    return &TableReservation{tables: t, users: u, payment: p}
}


func (uc *TableReservation) BookATable(ctx context.Context, _tableId string, _payment payment.DBPayment) (err error) {

	tableid, errParseInt := strconv.ParseInt(_tableId, 10, 8)
	if errParseInt != nil { return errParseInt }

	// ## valid cliente -> check
	fromUser, errUserNotFound := uc.users.FindUserById(ctx, _payment.ClientUUID)
	if errUserNotFound != nil { return errUserNotFound }

	// ## table reserved -> check
	available, errTableNotFound := uc.tables.IsAvailable(ctx, _tableId)
	if errTableNotFound != nil { return errTableNotFound }
	if !available {
		return fmt.Errorf("table is already reserved")
	}

	// ## save info on database
	//TODO: this
	// ideal:
	// 1. INSERT booking  → status: "pending"       [query 1]
	// 2. INSERT payment  → status: "pending", amount, currency, accounts  [query 2]
	//    // payment table existe desde el inicio, no al final
	// 3. Process payment → devuelve (transactionID, error)
	//    if error:
	//        UPDATE booking  → status: "payment_failed"   [query 3a]
	//        UPDATE payment  → status: "failed"            [query 3b]
	//        return ErrPaymentFailed  // limpio, sin huérfanos
	// 4. UPDATE booking  → status: "confirmed", paid: true  [query 3]
	// 5. UPDATE payment  → status: "completed", transaction_id: X  [query 4]

	bookingId, errBookingFailed := uc.reserv.Book(ctx,
		reservation.DBReservation{
			ClientUUID: fromUser.UUID,
			TableId:    int8(tableid),
			Paid:       true,
		},
	)

	if errBookingFailed  != nil {
		return errBookingFailed
	}

	_,errPaymentCreation := uc.paydb.SaveToDB(ctx, 
		payment.DBPayment{
			ReservationUUID: bookingId,
			ClientUUID: _payment.ClientUUID,
			Amount: _payment.Amount,
			Currency: _payment.Currency,
			FromAcc: _payment.FromAcc,
			ToAcc: _payment.ToAcc,
		},
	)

	if errPaymentCreation != nil {
		return errPaymentCreation 
	}

	// ## payment success -> check
	// NOTE: the underscore is the transaction info. may be needed
	_, errPaymentFailed := uc.payment.Pay(ctx, _payment)
	if errPaymentFailed != nil { 
		//TODO:
		// UPDATE booking  → status: "payment_failed"   [query 3a]
		// UPDATE payment  → status: "failed"            [query 3b]
		// return ErrPaymentFailed  // limpio, sin huérfanos
		return errPaymentFailed
	}

	// 4. UPDATE booking  → status: "confirmed", paid: true  [query 3]
	// 5. UPDATE payment  → status: "completed", transaction_id: X  [query 4]

	return nil
}

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

	// ## valid cliente -> check
	fromUser, errUserNotFound := uc.users.FindUserById(ctx, _payment.ClientUUID)
	if errUserNotFound != nil { return err }

	// ## table reserved -> check
	table, errTableNotFound := uc.tables.FindTableById(ctx, _tableId)
	if errTableNotFound != nil { return err }
	if table.Reserved {
		return fmt.Errorf("table is already reserved")
	}

	// ## payment success -> check
	//at this point the errors are recoverable
	info, errPaymentFailed := uc.payment.Pay(ctx, _payment)
	if errPaymentFailed != nil { return err }

	_ = payment.DBPayment{
		TransId:    info.TransId,
		CreatedAt:  info.CreatedAt,
		ClientUUID: fromUser.UUID,
		Amount:     info.TotalAmount,
		Currency:   info.Currency,
		FromAcc:    info.SenderAcc,
		ToAcc:      info.RecieverAcc,
	}

	// ## save info on database
	tableid, errParseInt := strconv.ParseInt(_tableId, 10, 8)
	if errParseInt != nil { return err }

	// at this point is where everything collapses
	//TODO: implemment this in a future:

	// actual flow (fragile):
	// 1. payment processing  ← no return point
	// 2. save information to db ← might fail

	// ideal:
	// 1. save info first into database with "pending payments"
	// 2. process payment
	// 3. if payment success, "confirmed", else retry with a goroutine or smth


	bookingId, errBookingFailed := uc.reserv.Book(ctx,
		reservation.DBReservation{
			ClientUUID: _payment.ClientUUID,
			Table:      int8(tableid),
			Paid:       true,
		},
	)

	//booking was unsuccessful - payment done. try to repay
	if errBookingFailed  != nil {
		_, errRefundFailed := uc.payment.Refund(ctx, _payment)
		if errRefundFailed   != nil {
			return fmt.Errorf("CRITICAL: payment charged but booking failed and refund failed: %v", errRefundFailed)
		}
		return errBookingFailed
	}
	
	errSaveFailed := uc.paydb.SaveToDB(ctx, _payment)
	if errSaveFailed != nil {
		cancelErr := uc.reserv.Cancel(ctx, bookingId)
		if cancelErr != nil {
			err = fmt.Errorf("%w (CRITICAL: reservation booked but cancellation failed: %v)", err, cancelErr)
		}
	}

	return nil
}

// i needed a scope to ensure all of database operations finish
// just fine. if one of them fails, the bloc where it's used
// will handle it.
func (uc *TableReservation) finalize(
	ctx context.Context,
	tableID string,
	body payment.DBPayment) (err error) {

	


	return nil
}

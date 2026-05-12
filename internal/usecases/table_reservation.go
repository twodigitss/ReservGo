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
	fromUser, err := uc.users.FindUserById(ctx, _payment.ClientUUID)
	if err != nil {
		return err
	}

	// ## table reserved -> check
	table, err := uc.tables.FindTableById(ctx, _tableId)
	if err != nil {
		return err
	}
	if table.Reserved {
		return fmt.Errorf("table is already reserved")
	}

	// ## payment success -> check
	info, err := uc.payment.Pay(ctx, _payment)
	if err != nil {
		return err
	}

	// Compensating transaction for payment: refund if later steps fail
	defer func() {
		if err != nil {
			_, refundErr := uc.payment.Refund(ctx, _payment)
			if refundErr != nil {
				// We wrap the original error to indicate that refund also failed
				err = fmt.Errorf("%w (CRITICAL: payment charged but refund failed: %v)", err, refundErr)
			}
		}
	}()

	saveto := payment.DBPayment{
		TransId:    info.TransId,
		CreatedAt:  info.CreatedAt,
		ClientUUID: fromUser.UUID,
		Amount:     info.TotalAmount,
		Currency:   info.Currency,
		FromAcc:    info.SenderAcc,
		ToAcc:      info.RecieverAcc,
	}

	// ## save info on database
	err = uc.finalize(ctx, _tableId, saveto)
	if err != nil {
		return err // The deferred refund will trigger here
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

	tableid, err := strconv.ParseInt(tableID, 10, 8)
	if err != nil {
		return err
	}

	reservationID, err := uc.reserv.Book(ctx,
		reservation.DBReservation{
			ClientUUID: body.ClientUUID,
			Table:      int8(tableid),
			Paid:       true,
		},
	)
	if err != nil {
		return err
	}

	// Compensating transaction: cancel reservation if later steps fail
	defer func() {
		if err != nil {
			cancelErr := uc.reserv.Cancel(ctx, reservationID)
			if cancelErr != nil {
				err = fmt.Errorf("%w (CRITICAL: reservation booked but cancellation failed: %v)", err, cancelErr)
			}
		}
	}()

	err = uc.paydb.SaveToDB(ctx, body)
	if err != nil {
		return err // The deferred cancellation will trigger here
	}

	return nil
}

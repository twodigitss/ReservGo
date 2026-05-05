package usecases

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/twodigitss/reserv-go/infrastructure/payment_provider"
	"github.com/twodigitss/reserv-go/internal/modules/tables"
	"github.com/twodigitss/reserv-go/internal/modules/users"
)

type TableReservation struct {
	tables *tables.TableService
	users  *users.UserService
}

func NewCreateReservation(t *tables.TableService, u *users.UserService) *TableReservation  {
    return &TableReservation {tables: t, users: u}
}


func (this *TableReservation ) BookATable(ctx context.Context, UUID string, TABLEID string, PAYMENT any) (error){
	/*
	insert table reservations_demo.reservations_online () values ()  returning &dbRow{}
	reservation row info →
	reservationId, created at, clientuuid, table, paid, visited, visited date.

	after that
		tengo el dbrow con la información.
		generar un qr o algo
	*/

	// ## valid cliente -> check
	user, err := this.users.FindUserById(ctx, UUID);
	if err != nil {

		if uuid.IsInvalidLengthError(err) {
			return fmt.Errorf("Invalid UUID length")
		}

		//might be not rows in result
		if err.Error() == "no rows in result" {
			return fmt.Errorf("No user was found")
		}
	}

	if err = uuid.Validate(user.UUID); err != nil {
		return fmt.Errorf("Invalid UUID")
	}

	// ## table reserved -> checl
	table, err := this.tables.FindTableById(ctx, TABLEID)
	if err != nil { return err }
	if table.Reserved {
		return fmt.Errorf("table is already reserved")
	}

	// ## payment success -> check
	sucess, err := paymentprovider.Pay(UUID, TABLEID)
	if err != nil { return err }
	fmt.Print(sucess)

	//insert in table
	//insert table reservations_demo.reservations_online () values () returning &dbRow{}
	// get &dbRow{}
	//generate qr or a confirmation it works

	return nil

}

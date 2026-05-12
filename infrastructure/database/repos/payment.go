package repos

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twodigitss/reserv-go/internal/modules/payment"
)

type PaymentImpl struct { DB *pgxpool.Pool }
func NewPaymentImpl(db *pgxpool.Pool) *PaymentImpl {
	return &PaymentImpl {DB: db}
}

func (this *PaymentImpl) SaveToDB(ctx context.Context, _body payment.DBPayment)(error){
	//FIX: i dont have a payment table xD

	//insert in table
	//insert table reservations_demo.reservations_online () values () returning &dbRow{}
	// get &dbRow{}

	//generate qr or a confirmation it works
	// _, err := this.DB.Exec(ctx,
	// 	`INSERT INTO reservations_demo.reservations_online 
	// 	(client_uuid, \"table\", paid) VALUES 
	// 	($1, $2, $3)`,
	// 	_body.ClientUUID, _body.ta, _body.Paid,
	// )
	// if err != nil { return err }
	return nil
}

func (this *PaymentImpl) GetByTransID(ctx context.Context, _id string)(error){
	return nil
}

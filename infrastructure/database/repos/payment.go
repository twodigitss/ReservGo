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

func (this *PaymentImpl) SaveToDB(ctx context.Context, _body payment.DBPayment)(payment.DBPayment, error){
	var res payment.DBPayment
	err := this.DB.QueryRow(ctx,
		`INSERT INTO dbv2.reservations
		(reservation_uuid, client_uuid, amount, currency, sender_acc, reciever_acc) VALUES 
		($1, $2, $3)`,
		_body.ReservationUUID, _body.ClientUUID, _body.Amount, 
		_body.Currency, _body.FromAcc, _body.ToAcc,
	).Scan(&res.ReservationUUID, &res.ClientUUID, &res.Amount, &res.Currency, &res.FromAcc, &res.ToAcc)

	if err != nil { return payment.DBPayment{}, err }
	return res, nil
}

func (this *PaymentImpl) GetByTransID(ctx context.Context, _uuid string)(payment.DBPayment, error){
	var res payment.DBPayment
	err := this.DB.QueryRow(ctx,
		`SELECT uuid, reservation_uuid, client_uuid, amount, currency, sender_acc, reciever_acc
		FROM dbv2.reservations
		WHERE uuid=$1`, _uuid,
	).Scan(&res.UUID ,&res.ReservationUUID, &res.ClientUUID, &res.Amount, &res.Currency, &res.FromAcc, &res.ToAcc)

	if err != nil { return payment.DBPayment{}, err }
	return res,nil
}

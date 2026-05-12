package reservation
import "context"

//se supone qee esta cosa es respecto a la tabla reservsiones de supa
type Reservation interface {
	Book(ctx context.Context, _body DBReservation) (string,error)
	Cancel(ctx context.Context, _id string) error
	Update(ctx context.Context, _id string) error
	Exists(ctx context.Context, _id string) (bool, error)
	GetByID(ctx context.Context, _id string) (DBReservation, error)
	GetByClientUUID(ctx context.Context, _uuid string) ([]DBReservation, error)
}

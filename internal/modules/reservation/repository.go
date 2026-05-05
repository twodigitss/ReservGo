package reservation
import "context"

//se supone qee esta cosa es respecto a la tabla reservsiones de supa
type Reservation interface {
	BookReservation(ctx context.Context, body HttpBodyReserv) error
}

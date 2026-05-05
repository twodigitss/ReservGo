package repos

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twodigitss/reserv-go/internal/modules/reservation"
)

type ReservRepoImpl struct { DB *pgxpool.Pool }

func NewReservRepo(db *pgxpool.Pool) *ReservRepoImpl {
	return &ReservRepoImpl {DB: db}
}

func (r *ReservRepoImpl) BookReservation(ctx context.Context, body reservation.HttpBodyReserv)  error {
	var result reservation.DBReservation

	// i should validate here that the paid has already been cleared
	if !body.Paid {
		return fmt.Errorf("Payment transaction has not been cleared yet, therefore reservation cannot be completed")
	}

	// que necesito?
	// client uuid
	// tabelid == INPUT
	// created id == default
	// client uuid = INPUT
	// paid = false default
	// paid == se supone que esta cosa no debería de estar
	// 	  rocesandose en el momento de hacer esta operación de reservation?
	// visited == false
	// visited date = default

	//why would i even need to return id/created_at?
	// should i return it?
	// for what do i even need the id/created_at
	r.DB.QueryRow(ctx,
		"INSERT INTO reservations_demo.reservations_online (client_uuid, table, paid) VALUES ($1, $2, $3) RETURNING id, created_at",
		body.ClientUUID, body.Table, body.Paid,
		//should i make the default values too?
	) .Scan(
		&result.ID,  &result.CreatedAt,
		&result.ClientUUID, &result.Table, &result.Paid, &result.Visited, &result.VisitedDate,
	)
	return nil

}

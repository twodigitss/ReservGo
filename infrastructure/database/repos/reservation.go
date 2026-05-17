package repos
import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twodigitss/reserv-go/internal/modules/reservation"
)

type ReservRepoImpl struct { DB *pgxpool.Pool }
func NewReservRepo(db *pgxpool.Pool) *ReservRepoImpl {
	return &ReservRepoImpl {DB: db}
}

// var is_this_file_implemmenting_correclty_the_interface reservation.Reservation = (*ReservRepoImpl)(nil)

func (this *ReservRepoImpl) Book(ctx context.Context, body reservation.DBReservation) (string,error){
	var reservationID string
	err := this.DB.QueryRow(ctx,
		`INSERT INTO dbv2.reservations 
		(client_uuid, table_id, paid) VALUES 
		($1, $2, $3) RETURNING uuid `, body.ClientUUID, body.TableId, body.Paid,
	).Scan(&reservationID)

	if err != nil { return "",err }
	return reservationID, nil
}

func (this *ReservRepoImpl) Cancel(ctx context.Context, _id string) error{
	_, err := this.DB.Exec(ctx, `DELETE FROM dbv2.reservations WHERE uuid = $1`, _id,)
	if err != nil { return err }
	return nil
}

func (this *ReservRepoImpl) Update(ctx context.Context, _id string) error{
	_, err := this.DB.Exec(ctx,
		`UPDATE dbv2.reservations 
		SET visited = TRUE 
		SET visited_date = Now() 
		WHERE uuid=$1`, _id,
	)
	if err != nil { return err }
	return nil
}

func (this *ReservRepoImpl) Exists(ctx context.Context, _id string) (bool, error){
	var exists bool
	err := this.DB.QueryRow(ctx,
		`SELECT EXISTS(SELECT 1 from dbv2.reservations WHERE uuid = $1)`, _id,
	).Scan(&exists)

	return exists, err
}

func (this *ReservRepoImpl) GetByID(ctx context.Context, _id string) (reservation.DBReservation, error){
	var res reservation.DBReservation
	err := this.DB.QueryRow(ctx,
		`SELECT uuid, client_uuid, table_id, created_at, paid, visited, visited_date
		 FROM dbv2.reservations 
		 WHERE uuid = $1`,
		_id,
	).Scan(&res.UUID, &res.ClientUUID, &res.TableId, &res.CreatedAt, &res.Paid, &res.Visited, &res.VisitedDate)

	if err != nil {
		return reservation.DBReservation{}, err
	}

	return res, nil
}

func (this *ReservRepoImpl) GetByClientUUID(ctx context.Context, _uuid string) ([]reservation.DBReservation, error){
	rows, err := this.DB.Query(ctx,
		`SELECT uuid, client_uuid, table_id, created_at, paid, visited, visited_date 
		 FROM dbv2.reservations 
		 WHERE client_uuid = $1`,
		_uuid,
	)

	if err != nil { return []reservation.DBReservation{}, err }
	defer rows.Close()

	var reservations []reservation.DBReservation
	for rows.Next() {
		var res reservation.DBReservation
		err := rows.Scan(&res.UUID, &res.ClientUUID, &res.TableId, &res.CreatedAt, &res.Paid, &res.Visited, &res.VisitedDate)

		if err != nil { return []reservation.DBReservation{}, err }
		reservations = append(reservations, res)
	}

	if err := rows.Err(); err != nil {
		return []reservation.DBReservation{}, err 
	}

	return reservations, nil
}

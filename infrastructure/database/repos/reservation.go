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

func (this *ReservRepoImpl) Book(ctx context.Context, body reservation.DBReservation) error{
	_, err := this.DB.Exec(ctx,
		`INSERT INTO reservations_demo.reservations_online 
		(client_uuid, \"table\", paid) VALUES 
		($1, $2, $3)`,
		body.ClientUUID, body.Table, body.Paid,
	)
	// might be a case where any database constraint triggers an error. 
	// will it handle itself through the exec func
	// or explicitly i will have to handle the error returned?
	if err != nil { return err }
	return nil

}

func (this *ReservRepoImpl) Cancel(ctx context.Context, _id string) error{
	_, err := this.DB.Exec(ctx,
		"DELETE FROM reservations_demo.reservations_online WHERE id = $1", _id,
	)
	if err != nil { return err }
	return nil
}

func (this *ReservRepoImpl) Update(ctx context.Context, _id string) error{
	_, err := this.DB.Exec(ctx,
		`UPDATE reservations_demo.reservations_online 
		SET visited = TRUE 
		WHERE id=$1`, _id,
	)
	if err != nil { return err }
	return nil
}

func (this *ReservRepoImpl) Exists(ctx context.Context, _id string) (bool, error){
	var exists bool
	this.DB.QueryRow(ctx,
		`SELECT EXISTS(SELECT 1 from reservations_demo.reservations_online 
		WHERE id = $1)`, _id,
	).Scan(exists)

	return exists, nil
}

func (this *ReservRepoImpl) GetByID(ctx context.Context, _id string) (*reservation.DBReservation, error){
	var res reservation.DBReservation
	err := this.DB.QueryRow(ctx,
		`SELECT * FROM reservations_demo.reservations_online 
		 WHERE id = $1 
		 RETURNING id,client_uuid,table,created_at,paid,visited,visited_date`,
		_id,
	).Scan(&res.ID, &res.ClientUUID, &res.Table, &res.CreatedAt, &res.Paid, &res.Visited, &res.VisitedDate)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (this *ReservRepoImpl) GetByClientUUID(ctx context.Context, _uuid string) (*[]reservation.DBReservation, error){
	rows, err := this.DB.Query(ctx,
		`SELECT id, client_uuid, table, created_at, paid, visited, visited_date 
		 FROM reservations_demo.reservations_online 
		 WHERE client_uuid = $1`,
		_uuid,
	)

	if err != nil { return nil, err }
	defer rows.Close()

	var reservations []reservation.DBReservation
	for rows.Next() {
		var res reservation.DBReservation
		err := rows.Scan(&res.ID, &res.ClientUUID, &res.Table, &res.CreatedAt, &res.Paid, &res.Visited, &res.VisitedDate)

		if err != nil { return nil, err }
		reservations = append(reservations, res)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &reservations, nil
}

package helpers

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// this function also works as a email finder, true if something was found
func DuplicatedEmail(ctx context.Context, DB *pgxpool.Pool, _email string) (bool, error) {
	var result bool
	err := DB.QueryRow(
		ctx, "SELECT EXISTS(SELECT 1 FROM reservations_demo.clients WHERE email=$1)", _email,
	).Scan(&result)
	if err != nil { return false, err }

	return result, nil

}

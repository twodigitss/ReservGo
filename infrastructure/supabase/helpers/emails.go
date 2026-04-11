package helpers

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// this function also works as a email finder, true if something was found
func DuplicatedEmail(ctx context.Context, DB *pgxpool.Pool, _email string) (bool, error) {
	var result string
	err := DB.QueryRow(ctx, "SELECT email FROM reservations_demo.clients WHERE email=$1", _email).Scan(&result)

	switch {
	case err != nil:
		return false, err
	case result == "":
		return false, nil
	default:
		return true, nil
	}

}

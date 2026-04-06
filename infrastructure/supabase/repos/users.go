package repos

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twodigitss/reserv-go/internal/modules/users"
)

// var is_this_file_implemmenting_correclty_the_interface users.UserModuleInterface= (*UserRepoImpl)(nil)

type UserRepoImpl struct { DB *pgxpool.Pool }

func NewUserRepo(db *pgxpool.Pool) *UserRepoImpl {
	return &UserRepoImpl {DB: db}
}

//Returns all the table users
func (this *UserRepoImpl) FindUsers(ctx context.Context) (*users.DBClients, error) {

	var result users.DBClients
	err := this.DB.QueryRow(
		context.Background(),
		"select uuid,created_at,name,last_name,email from reservations_demo.clients",
	).Scan(&result.UUID, &result.CreatedAt, &result.Name, &result.LastName, &result.Email)

	if err != nil {
		fmt.Print("ERROR GETTING PGX DATA: ", err)
		return nil, err
	}

	return (*users.DBClients)(&result), nil
}


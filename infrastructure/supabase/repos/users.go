package repos

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twodigitss/reserv-go/internal/modules/users"
	"github.com/twodigitss/reserv-go/infrastructure/supabase/helpers"
)

// var is_this_file_implemmenting_correclty_the_interface users.UserModuleInterface= (*UserRepoImpl)(nil)

type UserRepoImpl struct { DB *pgxpool.Pool }

func NewUserRepo(db *pgxpool.Pool) *UserRepoImpl {
	return &UserRepoImpl {DB: db}
}

//Returns all the table users
func (this *UserRepoImpl) ListUsers(ctx context.Context) ([]users.DBClient, error) {

	rows, err := this.DB.Query(
		ctx, "select uuid,created_at,name,last_name,email from reservations_demo.clients",
	)

	if err != nil {
		fmt.Print("ERROR GETTING QUERY: ", err)
		return nil, err
	}

	defer rows.Close()

	var user_ []users.DBClient

	for rows.Next(){
		var thing users.DBClient

		err := rows.Scan(&thing.UUID, &thing.CreatedAt, &thing.Name, &thing.LastName, &thing.Email,  )
		if err != nil {
			fmt.Print("ERROR SCANNING DATA", err)
			return nil, err
		}

		user_ = append(user_, thing)

	}

	if err:= rows.Err(); err!=nil{
		fmt.Print("ERROR DURING ROW ANALYSIS", err)
		return nil, err
	}

	return user_, nil
}

func (this *UserRepoImpl) FindUserById(ctx context.Context, _uuid string)(*users.DBClient, error){

	parsedID, err := uuid.Parse(_uuid) // id viene como string del path param
	if err != nil {
		fmt.Print("error parsing uuid")
		return nil, err
	}

	var result users.DBClient;
	err = this.DB.QueryRow(
		ctx, "SELECT uuid, name, last_name, email, created_at from reservations_demo.clients WHERE uuid=$1",
		parsedID,
	).Scan(&result.UUID, &result.Name, &result.LastName, &result.Email, &result.CreatedAt)

	if err != nil {
		fmt.Print("Error searching user", err)
		return nil, err
	}

	return &result, nil

}

func (this *UserRepoImpl) CreateUser(ctx context.Context, body users.DBClient) (*users.DBClient, error){

	duplicated, err := helpers.DuplicatedEmail(ctx, this.DB, body.Email)
	if duplicated {
		const res string = "User with same email already registered!"
		fmt.Println(res)
		return nil, fmt.Errorf(res)
	}
	if err != nil{
		fmt.Println("Error checking duplicated email", err)
		return nil, err
	}

	var query users.DBClient
	err = this.DB.QueryRow( ctx,
		`INSERT INTO reservations_demo.clients (name, last_name, email) values ($1,$2,$3)
		 RETURNING uuid, created_at; `, &body.Name, &body.LastName, &body.Email,
	).Scan(&query.UUID, &query.CreatedAt)

	if err != nil {
		fmt.Print("Error inserting user: ", err)
		return nil, err
	}

	query.Name = body.Name
	query.LastName = body.LastName
	query.Email = body.Email

	return &query, nil

}

func (this *UserRepoImpl) DeleteUser(ctx context.Context, _uuid string) (*users.DBClient, error) {

	parsedID, err := uuid.Parse(_uuid)
	if err != nil {
		fmt.Print("error parsing uuid")
		return nil, err
	}

	var query users.DBClient
	err = this.DB.QueryRow(
		ctx, `DELETE FROM reservations_demo.clients WHERE uuid = $1
		RETURNING uuid, email, name, last_name, created_at;`, &parsedID,
	).Scan(&query.UUID, &query.Email, &query.Name, &query.LastName, &query.CreatedAt)

	if err != nil {
		fmt.Print("Error deleting user", err.Error())
		return nil, err
	}

	return &query, nil

}

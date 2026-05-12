package users
import "time"

type DBClient struct {
	UUID 	  	string 		`db:"uuid"`
	CreatedAt time.Time `db:"created_at"`
	Name 	  	string 		`db:"name"`
	LastName  string 		`db:"last_name"`
	Email 	  string 		`db:"email"`
}

type HttpBodyClient struct {
	UUID 	  	string 		`json:"uuid" `
	CreatedAt time.Time `json:"created_at"`
	Name 	  	string 		`json:"name"`
	LastName  string 		`json:"last_name"`
	Email 	  string 		`json:"email" binding:"required,email"`
}

type HttpResClient struct {
	UUID 	  	string 		`json:"uuid" `
	CreatedAt time.Time `json:"created_at"`
	Name 	  	string 		`json:"name"`
	LastName  string 		`json:"last_name"`
	Email 	  string 		`json:"email"`
}

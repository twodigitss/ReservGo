package users
import "time"

type DBClients  struct {
	UUID string `db:"uuid"`
	CreatedAt time.Time `db:"created_at"`
	Name string `db:"name"`
	LastName string `db:"last_name"`
	Email string `db:"email"`
}

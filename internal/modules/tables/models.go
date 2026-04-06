package tables
import "time"

type DBTables  struct {
	Id string `db:"id"`
	TableName string `db:"table_name"`
	Reserved bool `db:"reserved"`
	CreatedAt time.Time `db:"created_at"`
}


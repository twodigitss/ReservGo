package tables
import "time"

type DBTables  struct {
	Id 		  string 	`db:"id"`
	CreatedAt time.Time `db:"created_at"`
	TableName string 	`db:"table_name"`
	Reserved  bool 	 	`db:"reserved"`
}

type HttpBodyTable struct {
	Id 		  string 	`json:"id"`
	CreatedAt time.Time `json:"created_at"`
	TableName string 	`json:"table_name"`
	Reserved  bool 		`json:"reserved"`
}

type HttpResTable struct {
	Id 		  string 	`json:"id"`
	CreatedAt time.Time `json:"created_at"`
	TableName string 	`json:"table_name"`
	Reserved  bool 	`json:"reserved"`
}

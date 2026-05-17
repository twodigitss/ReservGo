package tables
import "time"

type DBTables  struct {
	Id 		  	string 		`db:"id"`
	CreatedAt time.Time `db:"created_at"`
	Description string 	`db:"description"`
	Capacity string 		`db:"capacity"`
	Floor string 				`db:"floor"`
}

type HttpBodyTable struct {
	Id 		  	string 		`json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Description string 	`db:"description"`
	Capacity string 		`db:"capacity"`
	Floor string 				`db:"floor"`
}

type HttpResTable struct {
	Id 		  	string 		`json:"id"`
	Description string 	`db:"description"`
	Capacity string 		`db:"capacity"`
	Floor string 				`db:"floor"`
}

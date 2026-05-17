package reservation

import "time"

type DBReservation struct {
	UUID        string  	`db:"uuid"`
	ClientUUID  string    `db:"client_uuid"`
	TableId     int8   	  `db:"table_id"`
	Paid        bool      `db:"paid"`
	Visited     bool      `db:"visited"`
	VisitedDate time.Time `db:"visited_date"`
	CreatedAt 	time.Time `db:"created_at"`
}

type HttpBody struct {
	UUID        string  	`json:"uuid"`
	ClientUUID  string    `json:"client_uuid"`
	TableId     int8   	  `json:"table_id"`
	Paid        bool      `json:"paid"`
	Visited     bool      `json:"visited"`
	VisitedDate time.Time `json:"visited_date"`
	CreatedAt 	time.Time `json:"created_at"`

}

type HttpRes struct {
	UUID        string  	`json:"uuid"`
	ClientUUID  string    `json:"client_uuid"`
	TableId     int8   	  `json:"table_id"`
	Paid        bool      `json:"paid"`
	Visited     bool      `json:"visited"`
	VisitedDate time.Time `json:"visited_date"`
	CreatedAt 	time.Time `json:"created_at"`

}

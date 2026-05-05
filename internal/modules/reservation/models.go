package reservation

import "time"

type DBReservation struct {
	ID        string  	  `db:"id"`
	//esto es un "timestampz" en PostgreSQL con Z, no se si sea el mismo tipo qe time.time
	CreatedAt 	time.Time `db:"created_at"`
	ClientUUID  string    `db:"client_uuid"`
	Table       int8   	  `db:"table"`
	Paid        bool      `db:"paid"`
	Visited     bool      `db:"visited"`
	VisitedDate time.Time `db:"visited_date"`
}

type HttpBodyReserv struct {
	ID        string  	  `json:"id"`
	//esto es un "timestampz" en PostgreSQL con Z, no se si sea el mismo tipo qe time.time
	CreatedAt 	time.Time `json:"created_at"`
	ClientUUID  string    `json:"client_uuid"`
	Table       int8   	  `json:"table"`
	Paid        bool      `json:"paid"`
	Visited     bool      `json:"visited"`
	VisitedDate time.Time `json:"visited_date"`
}

type httpResReserv struct {
	ID        string  	  `json:"id"`
	//esto es un "timestampz" en PostgreSQL con Z, no se si sea el mismo tipo qe time.time
	CreatedAt 	time.Time `json:"created_at"`
	ClientUUID  string    `json:"client_uuid"`
	Table       int8   	  `json:"table"`
	Paid        bool      `json:"paid"`
	Visited     bool      `json:"visited"`
	VisitedDate time.Time `json:"visited_date"`

}

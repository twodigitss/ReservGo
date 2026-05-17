package payment

import (
	"time"
)

//this is used for the Payment Provider in infra/
type ProviderTransaction struct {
	TransId 			string 		// pi_1NXYZ1234567890abcdef
	CreatedAt 		time.Time // time.time
	Confirmation 	bool			// true - false
	TotalAmount 	float64		// 100.00
	Currency  		string    // mxn - usd - etc
	SenderAcc 		string		// cus_ABC1234567890
	RecieverAcc 	string		// cus_ABC1234567892
}

//this and below is used for the db provider infra/
type DBPayment struct { 
	UUID 							string  	`db:"uuid"`
	ReservationUUID  	string    `db:"reservation_uuid"`
	ClientUUID  			string    `db:"client_uuid"`
	Amount  					float64   `db:"amount"`
	Currency  				string    `db:"currency"`
	FromAcc 					string		`db:"sender_acc"`
	ToAcc 						string		`db:"reciever_acc"`
	CreatedAt 				time.Time `db:"created_at"`
}

type HttpBody struct { 
	UUID 							string  	`json:"uuid"`
	ReservationUUID  	string    `json:"reservation_uuid"`
	ClientUUID  			string    `json:"client_uuid"`
	Amount  					float64   `json:"amount"`
	Currency  				string    `json:"currency"`
	FromAcc 					string		`json:"sender_acc"`
}

type HttpRes struct { 
	UUID 							string  	`json:"uuid"`
	ReservationUUID  	string    `json:"reservation_uuid"`
	ClientUUID  			string    `json:"client_uuid"`
	Amount  					float64   `json:"amount"`
	Currency  				string    `json:"currency"`
	FromAcc 					string		`json:"sender_acc"`
	CreatedAt 				time.Time `json:"created_at"`
}

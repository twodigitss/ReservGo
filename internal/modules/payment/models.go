package payment

import (
	"time"
)

// WARN: the 'Amount' field in the database is 'int8', which only goes to 127.

type ProviderTransaction struct {
	TransId 			string 		// pi_1NXYZ1234567890abcdef
	CreatedAt 		time.Time // time.time
	Confirmation 	bool			// true - false
	TotalAmount 	float64		// 100.00
	Currency  		string    // mxn - usd - etc
	SenderAcc 		string		// cus_ABC1234567890
	RecieverAcc 	string		// cus_ABC1234567892
}

type DBPayment struct { 
	TransId 			string  	`db:"id"`
	CreatedAt 		time.Time `db:"created_at"`
	ClientUUID  	string    `db:"client_uuid"`
	Amount  			float64   `db:"amount"`
	Currency  		string    `db:"currency"`
	FromAcc 			string		`db:"from_account"`
	ToAcc 				string		`db:"to_account"`
}

type HttpBody struct { 
	TransId 			string  	`json:"id"`
	ClientUUID  	string    `json:"client_uuid"`
	Amount  			float64   `json:"amount"`
	Currency  		string    `json:"currency"`
	FromAcc 			string		`json:"from_account"`
}

type HttpRes struct { 
	TransId 			string  	`json:"id"`
	CreatedAt 		time.Time `json:"created_at"`
	ClientUUID  	string    `json:"client_uuid"`
	Amount  			float64   `json:"amount"`
	Currency  		string    `json:"currency"`
	FromAcc 			string		`json:"from_account"`
}

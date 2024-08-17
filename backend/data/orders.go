package data

type Order struct {
	BlikCode     string  `json:"blik_code"`
	BlikPassword string  `json:"blik_password"`
	Email        string  `json:"email"`
	FinalPrice   float64 `json:"final_price"`
	FullName     string  `json:"full_name"`
	ItemsCount   int     `json:"items_count"`
	PaczkomatID  string  `json:"paczkomat_id"`
	PhoneNumber  string  `json:"phone_number"`
	Products     string  `json:"products"`
}

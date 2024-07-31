package data

const PricesTable = `
create table if not exists prices (
	id integer primary autoincrement,
	capacity integer,
	strength integer,
	price decimal
)
`

type Price struct {
	Id       int     `json:"id"`
	Capacity int     `json:"capacity"`
	Strength int     `json:"strength"`
	Price    float64 `json:"price"`
}

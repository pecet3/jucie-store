package data

import (
	"database/sql"
	"fmt"
	"log"
)

const OrdersTable = `
create table if not exists orders (
    id integer primary key autoincrement,
    blik_code text,
    blik_password text,
    email text not null,
    final_price real not null,
    full_name text not null,
    items_count integer not null,
    paczkomat_id text,
    phone_number text,
    products text not null,
	is_done integer default 0
);`

type Order struct {
	Id           int     `json:"id"`
	BlikCode     string  `json:"blik_code" validate:"required,len=6"`
	BlikPassword string  `json:"blik_password" validate:"required,min=4,max=20"`
	Email        string  `json:"email" validate:"required,email"`
	FinalPrice   float64 `json:"final_price" validate:"required,gt=0"`
	FullName     string  `json:"full_name" validate:"required,max=60"`
	ItemsCount   int     `json:"items_count" validate:"required,gt=0"`
	PaczkomatID  string  `json:"paczkomat_id" validate:"required,max=10"`
	PhoneNumber  string  `json:"phone_number" validate:"required,max=16"`
	Products     string  `json:"products" validate:"required,max=10000"`
	IsDone       int     `json:"-"`
}

func (o Order) GetAll(db *sql.DB) ([]Order, error) {
	query := `
        SELECT id, blik_code, blik_password, email, final_price, full_name, items_count, paczkomat_id, phone_number, products, is_done
        FROM orders
    `
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying orders: %v", err)
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.Id,
			&order.BlikCode,
			&order.BlikPassword,
			&order.Email,
			&order.FinalPrice,
			&order.FullName,
			&order.ItemsCount,
			&order.PaczkomatID,
			&order.PhoneNumber,
			&order.Products,
			&order.IsDone,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning order row: %v", err)
		}
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating order rows: %v", err)
	}

	return orders, nil
}

func (o Order) GetById(db *sql.DB, id int) (*Order, error) {
	query := "SELECT id, blik_code, blik_password, email, final_price, full_name, items_count, paczkomat_id, phone_number, products, is_done FROM orders WHERE id = ?"
	row := db.QueryRow(query, id)
	var order Order
	err := row.Scan(&order.Id, &order.BlikCode, &order.BlikPassword, &order.Email, &order.FinalPrice, &order.FullName, &order.ItemsCount, &order.PaczkomatID, &order.PhoneNumber, &order.Products, &order.IsDone) // Add IsDone field here
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("order with id %d not found", id)
		}
		log.Println(err)
		return nil, err
	}
	return &order, nil
}

func (o Order) Add(db *sql.DB, order *Order) (int, error) {
	query := "INSERT INTO orders (blik_code, blik_password, email, final_price, full_name, items_count, paczkomat_id, phone_number, products, is_done) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"              // Include is_done field
	result, err := db.Exec(query, order.BlikCode, order.BlikPassword, order.Email, order.FinalPrice, order.FullName, order.ItemsCount, order.PaczkomatID, order.PhoneNumber, order.Products, order.IsDone) // Include is_done value
	if err != nil {
		log.Println(err)
		return -1, err
	}
	orderId, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return int(orderId), nil
}

func (o Order) RemoveById(db *sql.DB, id int) error {
	query := "DELETE FROM orders WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (o Order) Update(db *sql.DB, order *Order) error {
	query := "UPDATE orders SET blik_code = ?, blik_password = ?, email = ?, final_price = ?, full_name = ?, items_count = ?, paczkomat_id = ?, phone_number = ?, products = ?, is_done = ? WHERE id = ?"       // Include is_done field
	_, err := db.Exec(query, order.BlikCode, order.BlikPassword, order.Email, order.FinalPrice, order.FullName, order.ItemsCount, order.PaczkomatID, order.PhoneNumber, order.Products, order.IsDone, order.Id) // Include is_done value
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (o Order) UpdateIsDone(db *sql.DB, isDone bool, id int) error {
	var isDoneInt int
	if isDone {
		isDoneInt = 1
	}
	query := "update orders set is_done = ? where id = ?"
	result, err := db.Exec(query, isDoneInt, id)
	if err != nil {
		return err
	}
	log.Println("<Data> Update is Done in orders result: ", result)
	return nil
}

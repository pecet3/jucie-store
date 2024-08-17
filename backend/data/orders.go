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
    products text not null
);`

type Order struct {
	Id           int     `json:"id"`
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

func (o Order) GetAll(db *sql.DB) ([]Order, error) {
	query := `
        SELECT id, blik_code, blik_password, email, final_price, full_name, items_count, paczkomat_id, phone_number, products
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
	query := "SELECT id, blik_code, blik_password, email, final_price, full_name, items_count, paczkomat_id, phone_number, products FROM orders WHERE id = ?"
	row := db.QueryRow(query, id)
	var order Order
	err := row.Scan(&order.Id, &order.BlikCode, &order.BlikPassword, &order.Email, &order.FinalPrice, &order.FullName, &order.ItemsCount, &order.PaczkomatID, &order.PhoneNumber, &order.Products)
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
	query := "INSERT INTO orders (blik_code, blik_password, email, final_price, full_name, items_count, paczkomat_id, phone_number, products) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := db.Exec(query, order.BlikCode, order.BlikPassword, order.Email, order.FinalPrice, order.FullName, order.ItemsCount, order.PaczkomatID, order.PhoneNumber, order.Products)
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
	query := "UPDATE orders SET blik_code = ?, blik_password = ?, email = ?, final_price = ?, full_name = ?, items_count = ?, paczkomat_id = ?, phone_number = ?, products = ? WHERE id = ?"
	_, err := db.Exec(query, order.BlikCode, order.BlikPassword, order.Email, order.FinalPrice, order.FullName, order.ItemsCount, order.PaczkomatID, order.PhoneNumber, order.Products, order.Id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

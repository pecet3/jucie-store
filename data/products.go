package data

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const ProductsTable = `
create table if not exists products (
	id integer primary key autoincrement,
	name text not null,
	description text not null,
	image_url text default '',
	quantity int,
	price decimal default 0.0,
	category text not null,
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp
);
`

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (p Product) GetAll(db *sql.DB) ([]Product, error) {
	query := `
        SELECT id, name, description, image_url, quantity, price, category, created_at, updated_at
        FROM products
    `
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying products: %v", err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.ImageURL,
			&product.Quantity,
			&product.Price,
			&product.Category,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning product row: %v", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating product rows: %v", err)
	}

	return products, nil
}
func (p Product) GetById(db *sql.DB, id int) (*Product, error) {
	query := "select id, name, description, image_url, quantity, price, category, created_at, updated_at from products where id = ?"
	row := db.QueryRow(query, id)
	var product Product
	err := row.Scan(&product.Id, &product.Name, &product.Description, &product.ImageURL, &product.Quantity, &product.Price, &product.Category, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &product, nil
}

func (p Product) Add(db *sql.DB, name, description, imageURL string, quantity int, price float64, category string) (int, error) {
	query := "insert into products (name, description, image_url, quantity, price, category) values (?, ?, ?, ?, ?, ?)"
	result, err := db.Exec(query, name, description, imageURL, quantity, price, category)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	productId, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return int(productId), nil
}

func (p Product) RemoveById(db *sql.DB, id int) error {
	query := "delete from products where id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p Product) Edit(db *sql.DB, id int, name, description, imageURL string, quantity int, price float64, category string) error {
	query := "update products set name = ?, description = ?, image_url = ?, quantity = ?, price = ?, category = ?, updated_at = current_timestamp where id = ?"
	_, err := db.Exec(query, name, description, imageURL, quantity, price, category, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

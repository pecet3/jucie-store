package data

import (
	"database/sql"
	"fmt"
)

const PricesTable = `
create table if not exists prices (
    id integer primary key autoincrement,
    capacity integer,
    price real
)`

type Price struct {
	Id       int     `json:"id"`
	Capacity int     `json:"capacity"`
	Price    float64 `json:"price"`
}

func (pr Price) AddPrice(db *sql.DB, p Price) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO prices(capacity, price) VALUES(?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Capacity, p.Price)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (pr Price) DeletePrice(db *sql.DB, id int) error {
	stmt, err := db.Prepare("DELETE FROM prices WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

func (pr Price) UpdatePrice(db *sql.DB, p Price) error {
	stmt, err := db.Prepare("UPDATE prices SET price = ? WHERE capacity = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Price, p.Capacity)
	if err != nil {
		return err
	}

	return nil
}

func (pr Price) GetPrice(db *sql.DB, id int) (Price, error) {
	var p Price
	err := db.QueryRow("SELECT id, capacity, price FROM prices WHERE id = ?", id).Scan(&p.Id, &p.Capacity, &p.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, fmt.Errorf("price not found")
		}
		return p, err
	}
	return p, nil
}

func insertPrices(db *sql.DB) error {
	pricesSQL := `
    INSERT INTO prices (capacity, price) VALUES
    (30,  29.99),
    (60,  49.99),
    (100, 69.99),
    `

	_, err := db.Exec(pricesSQL)
	if err != nil {
		return fmt.Errorf("failed to insert prices: %v", err)
	}
	return nil
}

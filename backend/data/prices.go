package data

import (
	"database/sql"
	"fmt"
)

const PricesTable = `
create table if not exists prices (
    id integer primary key autoincrement,
    capacity integer,
    strength integer,
    price real
)`

type Price struct {
	Id       int     `json:"id"`
	Capacity int     `json:"capacity"`
	Strength int     `json:"strength"`
	Price    float64 `json:"price"`
}

func AddPrice(db *sql.DB, p Price) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO prices(capacity, strength, price) VALUES(?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Capacity, p.Strength, p.Price)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func DeletePrice(db *sql.DB, id int) error {
	stmt, err := db.Prepare("DELETE FROM prices WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

func UpdatePrice(db *sql.DB, p Price) error {
	stmt, err := db.Prepare("UPDATE prices SET price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Price, p.Id)
	if err != nil {
		return err
	}

	return nil
}

func GetPrice(db *sql.DB, id int) (Price, error) {
	var p Price
	err := db.QueryRow("SELECT id, capacity, strength, price FROM prices WHERE id = ?", id).Scan(&p.Id, &p.Capacity, &p.Strength, &p.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, fmt.Errorf("price not found")
		}
		return p, err
	}
	return p, nil
}

package data

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func newDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	prepareList := [2]string{
		PricesTable,
		ProductsTable,
	}

	for _, table := range prepareList {
		if err := prepare(db, table); err != nil {
			log.Fatalf("Failed to prepare DB: %v", err)
		}
	}
	insertPrices(db)
	log.Println("<DB> Prices have been inserted successfully")

	log.Println("<DB> Preparing DB has been finished")
	return db
}

func prepare(db *sql.DB, table string) error {
	s, err := db.Prepare(table)
	if err != nil {
		return err
	}
	_, err = s.Exec()
	if err != nil {
		return err
	}

	endIndex := strings.Index(table, "(")
	log.Print("<DB> ", table[1:endIndex])

	return nil
}

type Data struct {
	Db      *sql.DB
	Price   Price
	Product Product
}

func New() Data {
	return Data{
		Db:      newDb(),
		Price:   Price{},
		Product: Product{},
	}
}

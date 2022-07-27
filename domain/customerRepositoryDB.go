package domain

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct{}

// Create constructor
func NewCustomerRepositoryDB() CustomerRepositoryDB {
	return CustomerRepositoryDB{}
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	connStr := "postgres://postgres:ktl123@localhost/banking_cap?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	query := "select * from customers"

	rows, err := db.Query(query)
	if err != nil {
		log.Println("error query data to customer table", err.Error())
		return nil, err
	}

	var customers []Customer
	// Untuk proses setiap data yg masuk dari DB
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
		if err != nil {
			log.Println("error scanning customer data", err.Error())
		}
		customers = append(customers, c)
	}
	return customers, nil
}

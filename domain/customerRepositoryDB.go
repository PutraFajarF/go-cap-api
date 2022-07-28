package domain

import (
	"capi/errs"
	"capi/logger"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

// Create constructor
func NewCustomerRepositoryDB() CustomerRepositoryDB {
	// In development mode postgreSQL, sslmode setting to disable instead of verify-full
	connStr := "postgres://postgres:ktl123@localhost/banking_cap?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerRepositoryDB{db}
}

func (d CustomerRepositoryDB) FindByID(customerID string) (*Customer, *errs.AppErr) {
	query := "select * from customers where customer_id = $1"

	row := d.client.QueryRow(query, customerID)

	var c Customer
	err := row.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("error customer data not found" + err.Error())
			return nil, errs.NewNotFoundError("customer data not found")
		} else {
			logger.Error("error scanning customer data" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	// untuk balikin nilai struct pakai pointer
	return &c, nil
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {

	query := "select * from customers"

	rows, err := d.client.Query(query)
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
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

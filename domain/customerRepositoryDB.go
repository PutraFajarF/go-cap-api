package domain

import (
	"capi/errs"
	"capi/logger"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

// Create constructor
func NewCustomerRepositoryDB() CustomerRepositoryDB {
	// In development mode postgreSQL, sslmode setting to disable instead of verify-full
	connStr := "postgres://postgres:ktl123@localhost/banking_cap?sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerRepositoryDB{db}
}

func (d CustomerRepositoryDB) FindByID(customerID string) (*Customer, *errs.AppErr) {
	query := "select * from customers where customer_id = $1"

	var c Customer

	err := d.client.Get(&c, query, customerID)
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

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppErr) {
	var e []Customer

	if status == "" {
		query := "select * from customers"
		err := d.client.Select(&e, query)
		if err != nil {
			log.Println("error query data to customer table", err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	} else {
		if status == "active" {
			status = "1"
			query := "select * from customers where status = $1"
			err := d.client.Select(&e, query, status)
			if err != nil {
				log.Println("error query data to customer table", err.Error())
				return nil, errs.NewNotFoundError("customer data not found")
			}
		} else if status == "inactive" {
			status = "0"
			query := "select * from customers where status = $1"
			err := d.client.Select(&e, query, status)
			if err != nil {
				log.Println("error query data to customer table", err.Error())
				return nil, errs.NewNotFoundError("customer data not found")
			}
			// conditional statement jika query string mencari selain ?status=active atau ?status=inactive maka akan menghasilkan error 404 not found
		} else {
			query := "select * from customers where status = $1"
			err := d.client.Select(&e, query)
			if err != nil {
				log.Println("error query data to customer table", err.Error())
				return nil, errs.NewNotFoundError("customer data not found")
			}
		}
	}

	return e, nil
}

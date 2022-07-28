package domain

import "capi/errs"

type Customer struct {
	// untuk merubah key tampilan di json maupun xml dari prefix huruf besar menjadi huruf kecil
	ID          string `json:"id" xml:"id" db:"customer_id"`
	Name        string `json:"name" xml:"name"`
	City        string `json:"city" xml:"city"`
	ZipCode     string `json:"zip_code" xml:"zipcode"`
	DateOfBirth string `json:"date_of_birth" xml:"dateofbirth" db:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindByID(string) (*Customer, *errs.AppErr)
}

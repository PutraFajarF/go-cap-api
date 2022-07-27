package domain

// define struct
type CustomerRepositoryStub struct {
	Customer []Customer
}

// define method
func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.Customer, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "User1", "Jakarta", "12345", "2022-01-25", "1"},
		{"2", "User2", "Surabaya", "67890", "2022-02-09", "1"},
		{"3", "User3", "Semarang", "76584", "2022-03-07", "1"},
		{"4", "User4", "Bogor", "252567", "2022-05-08", "1"},
	}

	return CustomerRepositoryStub{Customer: customers}
}

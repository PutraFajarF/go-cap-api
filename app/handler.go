package app

import (
	"capi/service"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// var customers []Customer = []Customer{
// 	{1, "User1", "Jakarta", "12345"},
// 	{2, "User2", "Surabaya", "67890"},
// 	{3, "User3", "Semarang", "76584"},
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello Celerates!")
// }

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "get customer endpoint")

	customers, _ := ch.service.GetAllCustomer()
	// Request Header
	if r.Header.Get("Content-Type") == "application/xml" {
		// Response Header
		w.Header().Add("Content-Type", "application/xml")
		// Marshaling data structure to XML representation
		xml.NewEncoder(w).Encode(customers)
	} else {
		// Response Header
		w.Header().Add("Content-Type", "application/json")
		// Marshaling data structure to JSON Representation
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandler) getCustomerByID(w http.ResponseWriter, r *http.Request) {
	// Untuk mengambil ID, get route variable
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	customer, err := ch.service.GetCustomerByID(customerId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err.Error())
		return
	}

	// Convert string to int
	// id, err := strconv.Atoi(customerId)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprint(w, "invalid customer id")
	// 	return
	// }

	// Searching customer data
	// var cust Customer

	// for _, data := range customers {
	// 	if data.ID == id {
	// 		cust = data
	// 	}
	// }

	// if cust.ID == 0 {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	fmt.Fprint(w, "customer data not found")
	// 	return
	// }

	// Return customer data
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// func addCustomer(w http.ResponseWriter, r *http.Request) {
// 	// decode request body
// 	var cust Customer
// 	json.NewDecoder(r.Body).Decode(&cust)

// 	// generate new id
// 	nextID := getNextID()
// 	cust.ID = nextID

// 	// save data to array
// 	customers = append(customers, cust)
// 	w.WriteHeader(http.StatusCreated)
// 	fmt.Fprintln(w, "customer successfully created")
// }

// func getNextID() int {
// 	cust := customers[len(customers)-1]

// 	return cust.ID + 1
// }

// func updatedCustomer(w http.ResponseWriter, r *http.Request) {
// 	// Untuk mengambil ID, get route variable
// 	vars := mux.Vars(r)
// 	customerId, err := strconv.Atoi(vars["customer_id"])
// 	if err != nil {
// 		fmt.Println("Unable to convert to string")
// 	}

// 	// decode request body, parse data dari bodynya
// 	var updatedCustomer Customer
// 	json.NewDecoder(r.Body).Decode(&updatedCustomer)

// 	for i, c := range customers {
// 		if c.ID == customerId {
// 			customers = append(customers[:i], customers[i+1:]...)
// 			customers = append(customers, updatedCustomer)
// 		}
// 	}
// 	json.NewEncoder(w).Encode(customers)
// 	w.WriteHeader(http.StatusOK)
// }

// func deleteCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	customerId, err := strconv.Atoi(vars["customer_id"])
// 	if err != nil {
// 		fmt.Println("Unable to convert to string")
// 	}

// 	for i, c := range customers {
// 		if c.ID == customerId {
// 			customers = append(customers[:i], customers[i+1:]...)
// 		}
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(customers)
// }

package app

import (
	"capi/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	fmt.Println("status =>", status)
	customers, err := ch.service.GetAllCustomer(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, customers)
}

func (ch *CustomerHandler) getCustomerByID(w http.ResponseWriter, r *http.Request) {
	// Untuk mengambil ID, get route variable
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	customer, err := ch.service.GetCustomerByID(customerId)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	// Return customer data
	writeResponse(w, http.StatusOK, customer)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	// set Header terlebih dahulu baru writeHeadernya, jika terbalik WriteHeader akan lock header dengan setting text/plain
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
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

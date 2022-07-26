package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	// untuk merubah key tampilan di json maupun xml dari prefix huruf besar menjadi huruf kecil
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
}

var customers []Customer = []Customer{
	{"User1", "Jakarta", "12345"},
	{"User2", "Surabaya", "67890"},
	{"User3", "Semarang", "76584"},
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Celerates!")
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "get customer endpoint")
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

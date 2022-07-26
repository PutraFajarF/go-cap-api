package app

import "net/http"

func Start() {

	// * defining routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getCustomers)

	// * starting the server
	http.ListenAndServe(":8080", nil)
}

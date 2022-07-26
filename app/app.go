package app

import "net/http"

func Start() {
	mux := http.NewServeMux()
	// * defining routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getCustomers)

	// * starting the server
	http.ListenAndServe(":8080", mux)
}

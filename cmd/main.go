package cmd

import (
	"log"
	"net/http"
)

// TODO: load balance
func main() {
	http.HandleFunc("POST /logs", func(w http.ResponseWriter, r *http.Request) {
		// take body
		// write batch
		// return 201 only
	})

	// TODO: cache
	http.HandleFunc("GET /logs", func(w http.ResponseWriter, r *http.Request) {
		// create filters
		// fast query with filters
	})

	// TODO: generate a html page to view logs

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error when create http server: %v", err)
	}
}


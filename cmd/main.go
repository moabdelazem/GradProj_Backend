package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Setting Router For Navigation
	muxRouter := mux.NewRouter()

	// Handle the Route
	muxRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	// Serve The Server On Port 8080
	http.ListenAndServe(":8080", muxRouter)
}

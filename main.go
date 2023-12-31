package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type response struct {
	Message string `json:"message"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(response{Message: "Hello Not Obadah"})
		if err != nil {
			return
		}
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

var database = []Response{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			json.NewEncoder(w).Encode(database)
			return
		}

		if r.Method == "POST" {
			var newRes Response

			err := json.NewDecoder(r.Body).Decode(&newRes)
			if err != nil {
				http.Error(w, "JSON invalido", http.StatusBadRequest)
				return
			}

			database = append(database, newRes)

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newRes)

			return
		}
	})

	fmt.Println("Listening and Serving at localhost:8080")

	http.ListenAndServe(":8080", nil)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Blog struct {
	Blog string `json:"blog"`
}

func handleLogin(w http.ResponseWriter, r *http.Request) {

}

func handleSignup(w http.ResponseWriter, r *http.Request) {

}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data Blog
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Print or process the data
		fmt.Printf("Received blog: %s\n", data.Blog)
		fmt.Fprintf(w, " %s\n", data.Blog)
	}
}

func handleRetrieve(w http.ResponseWriter, r *http.Request) {

}

func handleUpdate(w http.ResponseWriter, r *http.Request) {

}

func handleDelete(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/save", handlePost)

	port := ":5000"
	fmt.Println("Server is running on port " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

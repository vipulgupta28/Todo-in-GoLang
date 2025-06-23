package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Struct to receive/send post data
type BlogPost struct {
	Content string `json:"content"`
}

// Handle saving a post
func savepostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var post BlogPost
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Append to db.txt
		f, err := os.OpenFile("db.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			http.Error(w, "Unable to write to file", http.StatusInternalServerError)
			return
		}
		defer f.Close()

		f.WriteString(post.Content + "\n")
		fmt.Fprint(w, "Saved!")
	}
}

// Handle getting all posts
func getPostsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("db.txt")
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Write(data)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/savepost", savepostHandler)
	http.HandleFunc("/getposts", getPostsHandler)

	fmt.Println("🚀 Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article is ...
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles is a list of Articles
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
		Article{Title: "Quavo", Desc: "White", Content: "COCO"},
	}
	fmt.Println("Endpoint hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}

// homepage displays "Homepage endpoint hit"
// /articles displays list of articles
func handlerRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handlerRequests()
}

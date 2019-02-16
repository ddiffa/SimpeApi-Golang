package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles [] Article

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "API Go", Desc: "Simple Api Golang", Content: "Hello world"},
		Article{Title: "Belajar Golang", Desc:"Get All Article", Content:"API"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", getAllArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
	handleRequest()
}

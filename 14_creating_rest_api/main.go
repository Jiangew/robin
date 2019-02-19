package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      int    `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

// Encoding our articles array into a JSON string and then writing as part of response
func retrieveArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	fmt.Println("Endpoint Hit: retrieveArticles")

	json.NewEncoder(w).Encode(articles)
}

func retrieveArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "Key: "+key)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", retrieveArticles)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleRequestsWithMux() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/v2", homePage)
	myRouter.HandleFunc("/v2/articles", retrieveArticles)
	myRouter.HandleFunc("/v2/article/{id}", retrieveArticle)

	log.Fatal(http.ListenAndServe(":9091", myRouter))
}

func main() {
	// handleRequests()

	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequestsWithMux()
}

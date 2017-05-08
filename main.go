package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/all", returnAllArticles)
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
    http.HandleFunc("/authors", returnAllAuthors)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    handleRequests()
}

type Article struct {
	Id int
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type Author struct {
	Name string `json:"Name"`
	Country string `json:Country`
}

type Articles []Article
type Authors []Author

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    articles := Articles{
        Article{Id: 1, Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: 2,Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }    
    fmt.Println("Endpoint Hit: returnAllArticles")
    
    json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]
    
    fmt.Fprintf(w, "Key: " + key)
}

func returnAllAuthors(w http.ResponseWriter, r *http.Request){
    authors := Authors{
        Author{Name: "John Doe", Country: "Kenya"},
        Author{Name: "Jane Doe", Country: "Ghana"},
    }    
    
    json.NewEncoder(w).Encode(authors)
}
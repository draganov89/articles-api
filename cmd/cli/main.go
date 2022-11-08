package main

import (
	"articles-api/internal/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/articles", handlers.HandleArticles)
	log.Println("Service starting... Listening on port 8092...")
	http.ListenAndServe(":8092", nil)
}

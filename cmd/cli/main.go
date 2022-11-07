package main

import (
	"log"
	"net/http"

	c "articles-api/internal/core"
)

func main() {

	api := c.NewArticlesAPI()

	articles := api.GetArticles()

	log.Println(articles)

	// http.HandleFunc("/hello", hello)
	// http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}

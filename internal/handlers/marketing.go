package handlers

import (
	c "articles-api/internal/core"
	"articles-api/internal/structs"
	"encoding/json"
	"log"
	"net/http"
)

func HandleArticles(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		log.Println("error: invalid request method used! Expected GET request!")
		resp := make(map[string]string)
		resp["error"] = "Invalid request method used! Expected GET request!"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
	log.Printf("[GET] %s\n", r.URL)

	articlesApi := c.NewArticlesAPI()
	adsApi := c.NewAdsAPI()

	articlesChan := make(chan []structs.Article)
	adsChan := make(chan []structs.Ad)

	go func() {
		articles, err := articlesApi.GetArticles()
		if err != nil {
			log.Printf("error getting articles: %s\n", err)
		}
		articlesChan <- articles
	}()

	go func() {
		ads, err := adsApi.GetAds()
		if err != nil {
			log.Printf("error getting ads: %s\n", err)
		}
		adsChan <- ads
	}()

	ads := <-adsChan
	log.Printf("%d ads retrieved\n", len(ads))
	articles := <-articlesChan
	log.Printf("%d articles retrieved\n", len(articles))

	respArr := c.PrepareMarketingData(articles, ads)

	jsonResp, err := json.Marshal(respArr)
	if err != nil {
		log.Fatalf("error in JSON marshal: %s\n", err)
	}
	w.Write(jsonResp)
}

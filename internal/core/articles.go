package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	s "articles-api/internal/structs"
)

type IArticlesAPI interface {
	GetArticles() ([]s.Article, error)
}

type ArticlesAPI struct {
}

func NewArticlesAPI() IArticlesAPI {
	return &ArticlesAPI{}
}

func (a *ArticlesAPI) GetArticles() ([]s.Article, error) {
	url := "https://storage.googleapis.com/aller-structure-task/articles.json"
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error getting data from %s\n", url)
		return nil, err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading body from GET request %s\n", url)
		return nil, err
	}

	var articles []s.Article

	err = json.Unmarshal(bytes, &articles)
	if err != nil {
		log.Printf("error unmarshaling request body %s\n", url)
		return nil, err
	}
	return articles, nil
}

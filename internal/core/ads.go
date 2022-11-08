package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	s "articles-api/internal/structs"
)

type IAdsAPI interface {
	GetAds() ([]s.Ad, error)
}

type AdsAPI struct {
}

func NewAdsAPI() IAdsAPI {
	return &AdsAPI{}
}

func (a *AdsAPI) GetAds() ([]s.Ad, error) {
	url := "https://storage.googleapis.com/aller-structure-task/contentmarketing.json"
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

	var response s.AdsResponseObject

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		log.Printf("error unmarshaling request body %s\n", url)
		return nil, err
	}
	return response.Response.Items, nil
}

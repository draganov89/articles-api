package core

import (
	s "articles-api/internal/structs"
)

const BATCH_SIZE int = 5

func PrepareMarketingData(articles []s.Article, ads []s.Ad) []interface{} {

	adsCount := len(articles) / BATCH_SIZE

	resultLen := adsCount + len(articles)

	result := make([]interface{}, 0, resultLen)

	offset := 0
	for {
		from := offset * BATCH_SIZE
		to := from + BATCH_SIZE
		if from >= len(articles) {
			break
		}
		if to > len(articles) {
			to = len(articles)
		}

		var article interface{}
		articlesBatch := articles[from:to]
		for _, val := range articlesBatch {
			article = val
			result = append(result, article)
		}

		if len(articlesBatch) < 5 {
			break
		}

		var ad interface{}
		if offset >= len(ads) {
			ad = getDefaultAd()
		} else {
			ad = ads[offset]
		}
		result = append(result, ad)

		offset++
	}
	return result

}

func getDefaultAd() interface{} {
	return struct {
		Type string
	}{
		Type: "Ad",
	}
}

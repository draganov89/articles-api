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
	for offset*BATCH_SIZE < len(articles) {
		from := offset * BATCH_SIZE
		to := from + BATCH_SIZE

		if to > len(articles) {
			to = len(articles)
		}

		articlesBatch := articles[from:to]
		for _, val := range articlesBatch {
			result = append(result, val)
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

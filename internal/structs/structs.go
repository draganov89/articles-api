package structs

type ArticlesResponseObject struct {
	HttpStatus int              `json:"httpStatus"`
	Response   ArticlesResponse `json:"response"`
}

type ArticlesResponse struct {
	Items []Article `json:"items"`
}

type Article struct {
	Type         string  `json:"type"`
	HarvesterId  string  `json:"harvesterId"`
	CerebroScore float64 `json:"cerebro-score"`
	Url          string  `json:"url"`
	Title        string  `json:"title"`
	CleanImage   string  `json:"cleanImage"`
}

// Could've done it with only one struct ...

type AdsResponseObject struct {
	HttpStatus int         `json:"httpStatus"`
	Response   AdsResponse `json:"response"`
}

type AdsResponse struct {
	Items []Ad `json:"items"`
}

type Ad struct {
	Type              string  `json:"type"`
	HarvesterId       string  `json:"harvesterId"`
	CommercialPartner string  `json:"commercialPartner"`
	LogoURL           string  `json:"logoURL"`
	CerebroScore      float64 `json:"cerebro-score"`
	Url               string  `json:"url"`
	Title             string  `json:"title"`
	CleanImage        string  `json:"cleanImage"`
}

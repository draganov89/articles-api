package structs

type Article struct {
	Type         string  `json:"type"`
	HarvesterId  string  `json:"harvesterId"`
	CerebroScore float64 `json:"cerebro-score"`
	Url          string  `json:"url"`
	Title        string  `json:"title"`
	CleanImage   string  `json:"cleanImage"`
}

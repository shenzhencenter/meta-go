package objects

type CatalogAdvertiserInsight struct {
	Category *string `json:"category,omitempty"`
	Country  *string `json:"country,omitempty"`
}

var CatalogAdvertiserInsightFields = struct {
	Category string
	Country  string
}{
	Category: "category",
	Country:  "country",
}

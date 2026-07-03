package objects

type ProductFeedRulesGet struct {
	Data    *[]map[string]interface{} `json:"data,omitempty"`
	Paging  *map[string]interface{}   `json:"paging,omitempty"`
	Summary *map[string]interface{}   `json:"summary,omitempty"`
}

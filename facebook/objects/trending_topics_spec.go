package objects

type TrendingTopicsSpec struct {
	IsAllTrending        *bool     `json:"is_all_trending,omitempty"`
	IsSpecialBudgetAlloc *bool     `json:"is_special_budget_alloc,omitempty"`
	TrendingTopics       *[]string `json:"trending_topics,omitempty"`
}

var TrendingTopicsSpecFields = struct {
	IsAllTrending        string
	IsSpecialBudgetAlloc string
	TrendingTopics       string
}{
	IsAllTrending:        "is_all_trending",
	IsSpecialBudgetAlloc: "is_special_budget_alloc",
	TrendingTopics:       "trending_topics",
}

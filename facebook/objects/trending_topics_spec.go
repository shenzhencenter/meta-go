package objects

type TrendingTopicsSpec struct {
	IsAllTrending        *bool     `json:"is_all_trending,omitempty"`
	IsSpecialBudgetAlloc *bool     `json:"is_special_budget_alloc,omitempty"`
	TrendingTopics       *[]string `json:"trending_topics,omitempty"`
}

package objects

type PageRecommendedPostsInfo struct {
	RecommendationSource *string                `json:"recommendation_source,omitempty"`
	RecommendedPosts     *[]RecommendedPagePost `json:"recommended_posts,omitempty"`
}

var PageRecommendedPostsInfoFields = struct {
	RecommendationSource string
	RecommendedPosts     string
}{
	RecommendationSource: "recommendation_source",
	RecommendedPosts:     "recommended_posts",
}

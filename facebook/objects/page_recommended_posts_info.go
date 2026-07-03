package objects

type PageRecommendedPostsInfo struct {
	RecommendationSource *string                `json:"recommendation_source,omitempty"`
	RecommendedPosts     *[]RecommendedPagePost `json:"recommended_posts,omitempty"`
}

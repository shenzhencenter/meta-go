package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Recommendation struct {
	CreatedTime        *core.Time              `json:"created_time,omitempty"`
	HasRating          *bool                   `json:"has_rating,omitempty"`
	HasReview          *bool                   `json:"has_review,omitempty"`
	OpenGraphStory     *map[string]interface{} `json:"open_graph_story,omitempty"`
	Rating             *int                    `json:"rating,omitempty"`
	RecommendationType *string                 `json:"recommendation_type,omitempty"`
	ReviewText         *string                 `json:"review_text,omitempty"`
	Reviewer           *User                   `json:"reviewer,omitempty"`
}

var RecommendationFields = struct {
	CreatedTime        string
	HasRating          string
	HasReview          string
	OpenGraphStory     string
	Rating             string
	RecommendationType string
	ReviewText         string
	Reviewer           string
}{
	CreatedTime:        "created_time",
	HasRating:          "has_rating",
	HasReview:          "has_review",
	OpenGraphStory:     "open_graph_story",
	Rating:             "rating",
	RecommendationType: "recommendation_type",
	ReviewText:         "review_text",
	Reviewer:           "reviewer",
}

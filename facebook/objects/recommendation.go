package objects

import (
	"time"
)

type Recommendation struct {
	CreatedTime        *time.Time              `json:"created_time,omitempty"`
	HasRating          *bool                   `json:"has_rating,omitempty"`
	HasReview          *bool                   `json:"has_review,omitempty"`
	OpenGraphStory     *map[string]interface{} `json:"open_graph_story,omitempty"`
	Rating             *int                    `json:"rating,omitempty"`
	RecommendationType *string                 `json:"recommendation_type,omitempty"`
	ReviewText         *string                 `json:"review_text,omitempty"`
	Reviewer           *User                   `json:"reviewer,omitempty"`
}

package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type RecommendedPagePost struct {
	IntentScore *float64 `json:"intent_score,omitempty"`
	IsIgMedia   *bool    `json:"is_ig_media,omitempty"`
	PostID      *core.ID `json:"post_id,omitempty"`
}

var RecommendedPagePostFields = struct {
	IntentScore string
	IsIgMedia   string
	PostID      string
}{
	IntentScore: "intent_score",
	IsIgMedia:   "is_ig_media",
	PostID:      "post_id",
}

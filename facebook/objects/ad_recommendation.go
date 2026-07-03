package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdRecommendation struct {
	BlameField         *string                           `json:"blame_field,omitempty"`
	Code               *int                              `json:"code,omitempty"`
	Confidence         *enums.AdrecommendationConfidence `json:"confidence,omitempty"`
	Importance         *enums.AdrecommendationImportance `json:"importance,omitempty"`
	Message            *string                           `json:"message,omitempty"`
	RecommendationData *AdRecommendationData             `json:"recommendation_data,omitempty"`
	Title              *string                           `json:"title,omitempty"`
	Value              *string                           `json:"value,omitempty"`
}

var AdRecommendationFields = struct {
	BlameField         string
	Code               string
	Confidence         string
	Importance         string
	Message            string
	RecommendationData string
	Title              string
	Value              string
}{
	BlameField:         "blame_field",
	Code:               "code",
	Confidence:         "confidence",
	Importance:         "importance",
	Message:            "message",
	RecommendationData: "recommendation_data",
	Title:              "title",
	Value:              "value",
}

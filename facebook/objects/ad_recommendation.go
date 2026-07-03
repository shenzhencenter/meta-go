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

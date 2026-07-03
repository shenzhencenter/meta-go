package enums

type AdrecommendationConfidence string

const (
	AdrecommendationConfidenceHigh   AdrecommendationConfidence = "HIGH"
	AdrecommendationConfidenceLow    AdrecommendationConfidence = "LOW"
	AdrecommendationConfidenceMedium AdrecommendationConfidence = "MEDIUM"
)

func (value AdrecommendationConfidence) String() string {
	return string(value)
}

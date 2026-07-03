package enums

type AdrecommendationImportance string

const (
	AdrecommendationImportanceHigh   AdrecommendationImportance = "HIGH"
	AdrecommendationImportanceLow    AdrecommendationImportance = "LOW"
	AdrecommendationImportanceMedium AdrecommendationImportance = "MEDIUM"
)

func (value AdrecommendationImportance) String() string {
	return string(value)
}

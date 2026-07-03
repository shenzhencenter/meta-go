package enums

type PagecopyrightManualClaimsActionEnumParam string

const (
	PagecopyrightManualClaimsActionEnumParamBlock           PagecopyrightManualClaimsActionEnumParam = "BLOCK"
	PagecopyrightManualClaimsActionEnumParamClaimAdEarnings PagecopyrightManualClaimsActionEnumParam = "CLAIM_AD_EARNINGS"
	PagecopyrightManualClaimsActionEnumParamManualReview    PagecopyrightManualClaimsActionEnumParam = "MANUAL_REVIEW"
	PagecopyrightManualClaimsActionEnumParamMonitor         PagecopyrightManualClaimsActionEnumParam = "MONITOR"
	PagecopyrightManualClaimsActionEnumParamRequestTakedown PagecopyrightManualClaimsActionEnumParam = "REQUEST_TAKEDOWN"
)

func (value PagecopyrightManualClaimsActionEnumParam) String() string {
	return string(value)
}

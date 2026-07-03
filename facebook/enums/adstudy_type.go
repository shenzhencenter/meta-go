package enums

type AdstudyType string

const (
	AdstudyTypeBackendAbTesting         AdstudyType = "BACKEND_AB_TESTING"
	AdstudyTypeContinuousLiftConfig     AdstudyType = "CONTINUOUS_LIFT_CONFIG"
	AdstudyTypeCreativeSpendEnforcement AdstudyType = "CREATIVE_SPEND_ENFORCEMENT"
	AdstudyTypeGeoLift                  AdstudyType = "GEO_LIFT"
	AdstudyTypeLift                     AdstudyType = "LIFT"
	AdstudyTypePortfolioOptimizer       AdstudyType = "PORTFOLIO_OPTIMIZER"
	AdstudyTypeSplitTest                AdstudyType = "SPLIT_TEST"
	AdstudyTypeVersionControl           AdstudyType = "VERSION_CONTROL"
)

func (value AdstudyType) String() string {
	return string(value)
}

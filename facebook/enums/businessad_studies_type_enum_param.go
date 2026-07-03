package enums

type BusinessadStudiesTypeEnumParam string

const (
	BusinessadStudiesTypeEnumParamBackendAbTesting         BusinessadStudiesTypeEnumParam = "BACKEND_AB_TESTING"
	BusinessadStudiesTypeEnumParamContinuousLiftConfig     BusinessadStudiesTypeEnumParam = "CONTINUOUS_LIFT_CONFIG"
	BusinessadStudiesTypeEnumParamCreativeSpendEnforcement BusinessadStudiesTypeEnumParam = "CREATIVE_SPEND_ENFORCEMENT"
	BusinessadStudiesTypeEnumParamGeoLift                  BusinessadStudiesTypeEnumParam = "GEO_LIFT"
	BusinessadStudiesTypeEnumParamLift                     BusinessadStudiesTypeEnumParam = "LIFT"
	BusinessadStudiesTypeEnumParamPortfolioOptimizer       BusinessadStudiesTypeEnumParam = "PORTFOLIO_OPTIMIZER"
	BusinessadStudiesTypeEnumParamSplitTest                BusinessadStudiesTypeEnumParam = "SPLIT_TEST"
	BusinessadStudiesTypeEnumParamVersionControl           BusinessadStudiesTypeEnumParam = "VERSION_CONTROL"
)

func (value BusinessadStudiesTypeEnumParam) String() string {
	return string(value)
}

package enums

type UseradStudiesTypeEnumParam string

const (
	UseradStudiesTypeEnumParamBackendAbTesting         UseradStudiesTypeEnumParam = "BACKEND_AB_TESTING"
	UseradStudiesTypeEnumParamContinuousLiftConfig     UseradStudiesTypeEnumParam = "CONTINUOUS_LIFT_CONFIG"
	UseradStudiesTypeEnumParamCreativeSpendEnforcement UseradStudiesTypeEnumParam = "CREATIVE_SPEND_ENFORCEMENT"
	UseradStudiesTypeEnumParamGeoLift                  UseradStudiesTypeEnumParam = "GEO_LIFT"
	UseradStudiesTypeEnumParamLift                     UseradStudiesTypeEnumParam = "LIFT"
	UseradStudiesTypeEnumParamPortfolioOptimizer       UseradStudiesTypeEnumParam = "PORTFOLIO_OPTIMIZER"
	UseradStudiesTypeEnumParamSplitTest                UseradStudiesTypeEnumParam = "SPLIT_TEST"
	UseradStudiesTypeEnumParamVersionControl           UseradStudiesTypeEnumParam = "VERSION_CONTROL"
)

func (value UseradStudiesTypeEnumParam) String() string {
	return string(value)
}

package enums

type AdaccountcampaignsbylabelsOperatorEnumParam string

const (
	AdaccountcampaignsbylabelsOperatorEnumParamAll AdaccountcampaignsbylabelsOperatorEnumParam = "ALL"
	AdaccountcampaignsbylabelsOperatorEnumParamAny AdaccountcampaignsbylabelsOperatorEnumParam = "ANY"
)

func (value AdaccountcampaignsbylabelsOperatorEnumParam) String() string {
	return string(value)
}

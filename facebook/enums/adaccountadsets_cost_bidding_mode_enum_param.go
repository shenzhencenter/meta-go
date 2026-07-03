package enums

type AdaccountadsetsCostBiddingModeEnumParam string

const (
	AdaccountadsetsCostBiddingModeEnumParamBalanced      AdaccountadsetsCostBiddingModeEnumParam = "BALANCED"
	AdaccountadsetsCostBiddingModeEnumParamCostFocused   AdaccountadsetsCostBiddingModeEnumParam = "COST_FOCUSED"
	AdaccountadsetsCostBiddingModeEnumParamVolumeFocused AdaccountadsetsCostBiddingModeEnumParam = "VOLUME_FOCUSED"
)

func (value AdaccountadsetsCostBiddingModeEnumParam) String() string {
	return string(value)
}

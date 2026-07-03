package enums

type AdcreativelinkdatamomentType string

const (
	AdcreativelinkdatamomentTypeFbLiveShopping AdcreativelinkdatamomentType = "FB_LIVE_SHOPPING"
	AdcreativelinkdatamomentTypeIgLiveShopping AdcreativelinkdatamomentType = "IG_LIVE_SHOPPING"
)

func (value AdcreativelinkdatamomentType) String() string {
	return string(value)
}

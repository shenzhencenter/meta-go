package enums

type AdcreativelinkdataimageoverlayspecTextType string

const (
	AdcreativelinkdataimageoverlayspecTextTypeAutomatedPersonalize AdcreativelinkdataimageoverlayspecTextType = "automated_personalize"
	AdcreativelinkdataimageoverlayspecTextTypeCustom               AdcreativelinkdataimageoverlayspecTextType = "custom"
	AdcreativelinkdataimageoverlayspecTextTypeDisclaimer           AdcreativelinkdataimageoverlayspecTextType = "disclaimer"
	AdcreativelinkdataimageoverlayspecTextTypeFromPrice            AdcreativelinkdataimageoverlayspecTextType = "from_price"
	AdcreativelinkdataimageoverlayspecTextTypeGuestRating          AdcreativelinkdataimageoverlayspecTextType = "guest_rating"
	AdcreativelinkdataimageoverlayspecTextTypePercentageOff        AdcreativelinkdataimageoverlayspecTextType = "percentage_off"
	AdcreativelinkdataimageoverlayspecTextTypePrice                AdcreativelinkdataimageoverlayspecTextType = "price"
	AdcreativelinkdataimageoverlayspecTextTypeStarRating           AdcreativelinkdataimageoverlayspecTextType = "star_rating"
	AdcreativelinkdataimageoverlayspecTextTypeStrikethroughPrice   AdcreativelinkdataimageoverlayspecTextType = "strikethrough_price"
	AdcreativelinkdataimageoverlayspecTextTypeSustainable          AdcreativelinkdataimageoverlayspecTextType = "sustainable"
)

func (value AdcreativelinkdataimageoverlayspecTextType) String() string {
	return string(value)
}

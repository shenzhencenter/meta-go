package enums

type AdcreativevideodatacustomoverlayspecBackgroundOpacity string

const (
	AdcreativevideodatacustomoverlayspecBackgroundOpacityHalf  AdcreativevideodatacustomoverlayspecBackgroundOpacity = "half"
	AdcreativevideodatacustomoverlayspecBackgroundOpacitySolid AdcreativevideodatacustomoverlayspecBackgroundOpacity = "solid"
)

func (value AdcreativevideodatacustomoverlayspecBackgroundOpacity) String() string {
	return string(value)
}

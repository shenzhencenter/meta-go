package enums

type AdcreativelinkdataimagelayerspecBlendingMode string

const (
	AdcreativelinkdataimagelayerspecBlendingModeLighten  AdcreativelinkdataimagelayerspecBlendingMode = "lighten"
	AdcreativelinkdataimagelayerspecBlendingModeMultiply AdcreativelinkdataimagelayerspecBlendingMode = "multiply"
	AdcreativelinkdataimagelayerspecBlendingModeNormal   AdcreativelinkdataimagelayerspecBlendingMode = "normal"
)

func (value AdcreativelinkdataimagelayerspecBlendingMode) String() string {
	return string(value)
}

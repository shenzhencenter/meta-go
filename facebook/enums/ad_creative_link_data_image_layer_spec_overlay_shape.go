package enums

type AdcreativelinkdataimagelayerspecOverlayShape string

const (
	AdcreativelinkdataimagelayerspecOverlayShapeCircle    AdcreativelinkdataimagelayerspecOverlayShape = "circle"
	AdcreativelinkdataimagelayerspecOverlayShapeNone      AdcreativelinkdataimagelayerspecOverlayShape = "none"
	AdcreativelinkdataimagelayerspecOverlayShapePill      AdcreativelinkdataimagelayerspecOverlayShape = "pill"
	AdcreativelinkdataimagelayerspecOverlayShapeRectangle AdcreativelinkdataimagelayerspecOverlayShape = "rectangle"
	AdcreativelinkdataimagelayerspecOverlayShapeTriangle  AdcreativelinkdataimagelayerspecOverlayShape = "triangle"
)

func (value AdcreativelinkdataimagelayerspecOverlayShape) String() string {
	return string(value)
}

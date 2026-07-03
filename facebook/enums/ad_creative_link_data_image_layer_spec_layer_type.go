package enums

type AdcreativelinkdataimagelayerspecLayerType string

const (
	AdcreativelinkdataimagelayerspecLayerTypeFrameOverlay AdcreativelinkdataimagelayerspecLayerType = "frame_overlay"
	AdcreativelinkdataimagelayerspecLayerTypeImage        AdcreativelinkdataimagelayerspecLayerType = "image"
	AdcreativelinkdataimagelayerspecLayerTypeTextOverlay  AdcreativelinkdataimagelayerspecLayerType = "text_overlay"
)

func (value AdcreativelinkdataimagelayerspecLayerType) String() string {
	return string(value)
}

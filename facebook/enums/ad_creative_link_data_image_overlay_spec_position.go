package enums

type AdcreativelinkdataimageoverlayspecPosition string

const (
	AdcreativelinkdataimageoverlayspecPositionBottomLeft  AdcreativelinkdataimageoverlayspecPosition = "bottom_left"
	AdcreativelinkdataimageoverlayspecPositionBottomRight AdcreativelinkdataimageoverlayspecPosition = "bottom_right"
	AdcreativelinkdataimageoverlayspecPositionTopLeft     AdcreativelinkdataimageoverlayspecPosition = "top_left"
	AdcreativelinkdataimageoverlayspecPositionTopRight    AdcreativelinkdataimageoverlayspecPosition = "top_right"
)

func (value AdcreativelinkdataimageoverlayspecPosition) String() string {
	return string(value)
}

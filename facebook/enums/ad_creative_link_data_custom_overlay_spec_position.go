package enums

type AdcreativelinkdatacustomoverlayspecPosition string

const (
	AdcreativelinkdatacustomoverlayspecPositionBottomLeft  AdcreativelinkdatacustomoverlayspecPosition = "bottom_left"
	AdcreativelinkdatacustomoverlayspecPositionBottomRight AdcreativelinkdatacustomoverlayspecPosition = "bottom_right"
	AdcreativelinkdatacustomoverlayspecPositionTopLeft     AdcreativelinkdatacustomoverlayspecPosition = "top_left"
	AdcreativelinkdatacustomoverlayspecPositionTopRight    AdcreativelinkdatacustomoverlayspecPosition = "top_right"
)

func (value AdcreativelinkdatacustomoverlayspecPosition) String() string {
	return string(value)
}

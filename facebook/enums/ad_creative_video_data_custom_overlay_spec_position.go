package enums

type AdcreativevideodatacustomoverlayspecPosition string

const (
	AdcreativevideodatacustomoverlayspecPositionMiddleCenter AdcreativevideodatacustomoverlayspecPosition = "middle_center"
	AdcreativevideodatacustomoverlayspecPositionMiddleLeft   AdcreativevideodatacustomoverlayspecPosition = "middle_left"
	AdcreativevideodatacustomoverlayspecPositionMiddleRight  AdcreativevideodatacustomoverlayspecPosition = "middle_right"
	AdcreativevideodatacustomoverlayspecPositionTopCenter    AdcreativevideodatacustomoverlayspecPosition = "top_center"
	AdcreativevideodatacustomoverlayspecPositionTopLeft      AdcreativevideodatacustomoverlayspecPosition = "top_left"
	AdcreativevideodatacustomoverlayspecPositionTopRight     AdcreativevideodatacustomoverlayspecPosition = "top_right"
)

func (value AdcreativevideodatacustomoverlayspecPosition) String() string {
	return string(value)
}

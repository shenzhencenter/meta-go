package enums

type AdcreativelinkdataimageoverlayspecOverlayTemplate string

const (
	AdcreativelinkdataimageoverlayspecOverlayTemplateCircleWithText   AdcreativelinkdataimageoverlayspecOverlayTemplate = "circle_with_text"
	AdcreativelinkdataimageoverlayspecOverlayTemplatePillWithText     AdcreativelinkdataimageoverlayspecOverlayTemplate = "pill_with_text"
	AdcreativelinkdataimageoverlayspecOverlayTemplateTriangleWithText AdcreativelinkdataimageoverlayspecOverlayTemplate = "triangle_with_text"
)

func (value AdcreativelinkdataimageoverlayspecOverlayTemplate) String() string {
	return string(value)
}

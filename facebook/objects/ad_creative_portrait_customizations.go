package objects

type AdCreativePortraitCustomizations struct {
	CarouselDeliveryMode *string                   `json:"carousel_delivery_mode,omitempty"`
	Specifications       *[]map[string]interface{} `json:"specifications,omitempty"`
}

var AdCreativePortraitCustomizationsFields = struct {
	CarouselDeliveryMode string
	Specifications       string
}{
	CarouselDeliveryMode: "carousel_delivery_mode",
	Specifications:       "specifications",
}

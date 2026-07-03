package objects

type AdCreativePortraitCustomizations struct {
	CarouselDeliveryMode *string                   `json:"carousel_delivery_mode,omitempty"`
	Specifications       *[]map[string]interface{} `json:"specifications,omitempty"`
}

package objects

type ProductItemShipping struct {
	ShippingCountry       *string  `json:"shipping_country,omitempty"`
	ShippingPriceCurrency *string  `json:"shipping_price_currency,omitempty"`
	ShippingPriceValue    *float64 `json:"shipping_price_value,omitempty"`
	ShippingRegion        *string  `json:"shipping_region,omitempty"`
	ShippingService       *string  `json:"shipping_service,omitempty"`
}

var ProductItemShippingFields = struct {
	ShippingCountry       string
	ShippingPriceCurrency string
	ShippingPriceValue    string
	ShippingRegion        string
	ShippingService       string
}{
	ShippingCountry:       "shipping_country",
	ShippingPriceCurrency: "shipping_price_currency",
	ShippingPriceValue:    "shipping_price_value",
	ShippingRegion:        "shipping_region",
	ShippingService:       "shipping_service",
}

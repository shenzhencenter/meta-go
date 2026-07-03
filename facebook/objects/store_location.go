package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type StoreLocation struct {
	FullAddress   *string                 `json:"full_address,omitempty"`
	Hours         *map[string]interface{} `json:"hours,omitempty"`
	ID            *core.ID                `json:"id,omitempty"`
	PhoneNumber   *string                 `json:"phone_number,omitempty"`
	PickupOptions *[]string               `json:"pickup_options,omitempty"`
	PriceRange    *string                 `json:"price_range,omitempty"`
	StoreCode     *string                 `json:"store_code,omitempty"`
	ZipCode       *string                 `json:"zip_code,omitempty"`
}

var StoreLocationFields = struct {
	FullAddress   string
	Hours         string
	ID            string
	PhoneNumber   string
	PickupOptions string
	PriceRange    string
	StoreCode     string
	ZipCode       string
}{
	FullAddress:   "full_address",
	Hours:         "hours",
	ID:            "id",
	PhoneNumber:   "phone_number",
	PickupOptions: "pickup_options",
	PriceRange:    "price_range",
	StoreCode:     "store_code",
	ZipCode:       "zip_code",
}

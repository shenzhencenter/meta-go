package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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

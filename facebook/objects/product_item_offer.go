package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductItemOffer struct {
	AvailabilityArea   *[]map[string]interface{} `json:"availability_area,omitempty"`
	AvailabilityRadius *float64                  `json:"availability_radius,omitempty"`
	ID                 *core.ID                  `json:"id,omitempty"`
}

var ProductItemOfferFields = struct {
	AvailabilityArea   string
	AvailabilityRadius string
	ID                 string
}{
	AvailabilityArea:   "availability_area",
	AvailabilityRadius: "availability_radius",
	ID:                 "id",
}

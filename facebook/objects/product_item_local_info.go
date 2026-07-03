package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductItemLocalInfo struct {
	AvailabilityCircleOrigin       *ProductItemLocalInfoLatLongShape   `json:"availability_circle_origin,omitempty"`
	AvailabilityCircleRadius       *float64                            `json:"availability_circle_radius,omitempty"`
	AvailabilityCircleRadiusUnit   *string                             `json:"availability_circle_radius_unit,omitempty"`
	AvailabilityPolygonCoordinates *[]ProductItemLocalInfoLatLongShape `json:"availability_polygon_coordinates,omitempty"`
	AvailabilityPostalCodes        *[]string                           `json:"availability_postal_codes,omitempty"`
	AvailabilitySource             *string                             `json:"availability_source,omitempty"`
	ID                             *core.ID                            `json:"id,omitempty"`
	InferredCircleOrigin           *ProductItemLocalInfoLatLongShape   `json:"inferred_circle_origin,omitempty"`
	InferredCircleRadius           *float64                            `json:"inferred_circle_radius,omitempty"`
}

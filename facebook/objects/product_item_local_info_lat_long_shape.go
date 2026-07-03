package objects

type ProductItemLocalInfoLatLongShape struct {
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
}

var ProductItemLocalInfoLatLongShapeFields = struct {
	Latitude  string
	Longitude string
}{
	Latitude:  "latitude",
	Longitude: "longitude",
}

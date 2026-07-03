package objects

type CatalogBasedTargeting struct {
	GeoTargetingType *string `json:"geo_targeting_type,omitempty"`
}

var CatalogBasedTargetingFields = struct {
	GeoTargetingType string
}{
	GeoTargetingType: "geo_targeting_type",
}

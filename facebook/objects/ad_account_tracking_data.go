package objects

type AdAccountTrackingData struct {
	TrackingSpecs *ConversionActionQuery `json:"tracking_specs,omitempty"`
}

var AdAccountTrackingDataFields = struct {
	TrackingSpecs string
}{
	TrackingSpecs: "tracking_specs",
}

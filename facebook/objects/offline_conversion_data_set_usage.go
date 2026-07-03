package objects

type OfflineConversionDataSetUsage struct {
	NumLiftStudies *int `json:"num_lift_studies,omitempty"`
}

var OfflineConversionDataSetUsageFields = struct {
	NumLiftStudies string
}{
	NumLiftStudies: "num_lift_studies",
}

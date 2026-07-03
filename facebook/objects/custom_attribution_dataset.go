package objects

type CustomAttributionDataset struct {
	App   *[]map[string]interface{} `json:"app,omitempty"`
	Pixel *[]map[string]interface{} `json:"pixel,omitempty"`
}

var CustomAttributionDatasetFields = struct {
	App   string
	Pixel string
}{
	App:   "app",
	Pixel: "pixel",
}

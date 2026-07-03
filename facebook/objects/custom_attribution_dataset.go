package objects

type CustomAttributionDataset struct {
	App   *[]map[string]interface{} `json:"app,omitempty"`
	Pixel *[]map[string]interface{} `json:"pixel,omitempty"`
}

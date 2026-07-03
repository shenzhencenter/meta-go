package objects

type AdsTabular struct {
	Rows *[]map[string]interface{} `json:"rows,omitempty"`
}

package objects

type AdsTabular struct {
	Rows *[]map[string]interface{} `json:"rows,omitempty"`
}

var AdsTabularFields = struct {
	Rows string
}{
	Rows: "rows",
}

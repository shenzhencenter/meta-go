package objects

type AdCreativeInsights struct {
	Aesthetics *[]map[string]string `json:"aesthetics,omitempty"`
}

var AdCreativeInsightsFields = struct {
	Aesthetics string
}{
	Aesthetics: "aesthetics",
}

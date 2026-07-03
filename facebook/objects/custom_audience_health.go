package objects

type CustomAudienceHealth struct {
	Health *[]map[string]interface{} `json:"health,omitempty"`
}

var CustomAudienceHealthFields = struct {
	Health string
}{
	Health: "health",
}

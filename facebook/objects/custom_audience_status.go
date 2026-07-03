package objects

type CustomAudienceStatus struct {
	Code        *uint64 `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

var CustomAudienceStatusFields = struct {
	Code        string
	Description string
}{
	Code:        "code",
	Description: "description",
}

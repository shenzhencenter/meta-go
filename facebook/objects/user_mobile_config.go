package objects

type UserMobileConfig struct {
	SectionName *string                 `json:"section_name,omitempty"`
	Value       *map[string]interface{} `json:"value,omitempty"`
}

var UserMobileConfigFields = struct {
	SectionName string
	Value       string
}{
	SectionName: "section_name",
	Value:       "value",
}

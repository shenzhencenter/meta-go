package objects

type UserMobileConfig struct {
	SectionName *string                 `json:"section_name,omitempty"`
	Value       *map[string]interface{} `json:"value,omitempty"`
}

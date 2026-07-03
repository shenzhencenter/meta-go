package objects

type AdCreativePlatformCustomization struct {
	Instagram *map[string]interface{} `json:"instagram,omitempty"`
}

var AdCreativePlatformCustomizationFields = struct {
	Instagram string
}{
	Instagram: "instagram",
}

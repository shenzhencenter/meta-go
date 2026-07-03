package objects

type AdCreativeGenerativeAssetSpec struct {
	TransparencyMetadata *map[string]interface{} `json:"transparency_metadata,omitempty"`
}

var AdCreativeGenerativeAssetSpecFields = struct {
	TransparencyMetadata string
}{
	TransparencyMetadata: "transparency_metadata",
}

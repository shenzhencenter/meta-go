package objects

type CreativeAssetTag struct {
	Name *string `json:"name,omitempty"`
}

var CreativeAssetTagFields = struct {
	Name string
}{
	Name: "name",
}

package objects

type ProductVisualVariant struct {
	Items     *[]ProductItem `json:"items,omitempty"`
	UniqueKey *string        `json:"unique_key,omitempty"`
}

var ProductVisualVariantFields = struct {
	Items     string
	UniqueKey string
}{
	Items:     "items",
	UniqueKey: "unique_key",
}

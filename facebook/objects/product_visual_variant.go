package objects

type ProductVisualVariant struct {
	Items     *[]ProductItem `json:"items,omitempty"`
	UniqueKey *string        `json:"unique_key,omitempty"`
}

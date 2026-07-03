package objects

type ProductVariant struct {
	Label        *string   `json:"label,omitempty"`
	Options      *[]string `json:"options,omitempty"`
	ProductField *string   `json:"product_field,omitempty"`
}

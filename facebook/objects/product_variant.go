package objects

type ProductVariant struct {
	Label        *string   `json:"label,omitempty"`
	Options      *[]string `json:"options,omitempty"`
	ProductField *string   `json:"product_field,omitempty"`
}

var ProductVariantFields = struct {
	Label        string
	Options      string
	ProductField string
}{
	Label:        "label",
	Options:      "options",
	ProductField: "product_field",
}

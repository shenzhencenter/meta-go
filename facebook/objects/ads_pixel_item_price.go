package objects

type AdsPixelItemPrice struct {
	Date              *string `json:"date,omitempty"`
	ItemPriceCoverage *string `json:"item_price_coverage,omitempty"`
}

var AdsPixelItemPriceFields = struct {
	Date              string
	ItemPriceCoverage string
}{
	Date:              "date",
	ItemPriceCoverage: "item_price_coverage",
}

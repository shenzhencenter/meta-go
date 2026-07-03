package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductSet struct {
	AutoCreationURL *string             `json:"auto_creation_url,omitempty"`
	Filter          *string             `json:"filter,omitempty"`
	ID              *core.ID            `json:"id,omitempty"`
	LatestMetadata  *ProductSetMetadata `json:"latest_metadata,omitempty"`
	LiveMetadata    *ProductSetMetadata `json:"live_metadata,omitempty"`
	Name            *string             `json:"name,omitempty"`
	OrderingInfo    *[]int              `json:"ordering_info,omitempty"`
	ProductCatalog  *ProductCatalog     `json:"product_catalog,omitempty"`
	ProductCount    *uint64             `json:"product_count,omitempty"`
	RetailerID      *core.ID            `json:"retailer_id,omitempty"`
}

var ProductSetFields = struct {
	AutoCreationURL string
	Filter          string
	ID              string
	LatestMetadata  string
	LiveMetadata    string
	Name            string
	OrderingInfo    string
	ProductCatalog  string
	ProductCount    string
	RetailerID      string
}{
	AutoCreationURL: "auto_creation_url",
	Filter:          "filter",
	ID:              "id",
	LatestMetadata:  "latest_metadata",
	LiveMetadata:    "live_metadata",
	Name:            "name",
	OrderingInfo:    "ordering_info",
	ProductCatalog:  "product_catalog",
	ProductCount:    "product_count",
	RetailerID:      "retailer_id",
}

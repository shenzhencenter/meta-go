package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type DynamicItemDisplayBundleFolder struct {
	CategorizationCriteria *string                `json:"categorization_criteria,omitempty"`
	ID                     *core.ID               `json:"id,omitempty"`
	Name                   *string                `json:"name,omitempty"`
	ProductCatalog         *ProductCatalog        `json:"product_catalog,omitempty"`
	ProductSet             *ProductSet            `json:"product_set,omitempty"`
	ValidLabels            *[]map[string][]string `json:"valid_labels,omitempty"`
}

var DynamicItemDisplayBundleFolderFields = struct {
	CategorizationCriteria string
	ID                     string
	Name                   string
	ProductCatalog         string
	ProductSet             string
	ValidLabels            string
}{
	CategorizationCriteria: "categorization_criteria",
	ID:                     "id",
	Name:                   "name",
	ProductCatalog:         "product_catalog",
	ProductSet:             "product_set",
	ValidLabels:            "valid_labels",
}

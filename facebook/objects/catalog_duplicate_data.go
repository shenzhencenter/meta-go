package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CatalogDuplicateData struct {
	BestCatalogID       *core.ID `json:"best_catalog_id,omitempty"`
	IsDuplicatedCatalog *bool    `json:"is_duplicated_catalog,omitempty"`
}

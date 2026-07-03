package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CatalogDuplicateData struct {
	BestCatalogID       *core.ID `json:"best_catalog_id,omitempty"`
	IsDuplicatedCatalog *bool    `json:"is_duplicated_catalog,omitempty"`
}

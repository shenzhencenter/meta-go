package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductCatalogDataSource struct {
	AppID               *core.ID `json:"app_id,omitempty"`
	ID                  *core.ID `json:"id,omitempty"`
	IngestionSourceType *string  `json:"ingestion_source_type,omitempty"`
	Name                *string  `json:"name,omitempty"`
	UploadType          *string  `json:"upload_type,omitempty"`
}

var ProductCatalogDataSourceFields = struct {
	AppID               string
	ID                  string
	IngestionSourceType string
	Name                string
	UploadType          string
}{
	AppID:               "app_id",
	ID:                  "id",
	IngestionSourceType: "ingestion_source_type",
	Name:                "name",
	UploadType:          "upload_type",
}

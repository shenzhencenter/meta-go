package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessObject struct {
	Asset     *map[string]interface{} `json:"asset,omitempty"`
	AssetType *string                 `json:"asset_type,omitempty"`
	ID        *core.ID                `json:"id,omitempty"`
	Name      *string                 `json:"name,omitempty"`
	Picture   *string                 `json:"picture,omitempty"`
}

var BusinessObjectFields = struct {
	Asset     string
	AssetType string
	ID        string
	Name      string
	Picture   string
}{
	Asset:     "asset",
	AssetType: "asset_type",
	ID:        "id",
	Name:      "name",
	Picture:   "picture",
}

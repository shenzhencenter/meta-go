package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CatalogContentVersionConfig struct {
	ID      *core.ID `json:"id,omitempty"`
	Name    *string  `json:"name,omitempty"`
	Version *string  `json:"version,omitempty"`
}

var CatalogContentVersionConfigFields = struct {
	ID      string
	Name    string
	Version string
}{
	ID:      "id",
	Name:    "name",
	Version: "version",
}

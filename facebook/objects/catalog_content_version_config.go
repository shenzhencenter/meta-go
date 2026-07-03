package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CatalogContentVersionConfig struct {
	ID      *core.ID `json:"id,omitempty"`
	Name    *string  `json:"name,omitempty"`
	Version *string  `json:"version,omitempty"`
}

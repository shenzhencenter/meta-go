package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AppPublisher struct {
	ContentID *core.ID `json:"content_id,omitempty"`
	IconURL   *string  `json:"icon_url,omitempty"`
	ID        *core.ID `json:"id,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Platform  *string  `json:"platform,omitempty"`
	StoreName *string  `json:"store_name,omitempty"`
	StoreURL  *string  `json:"store_url,omitempty"`
}

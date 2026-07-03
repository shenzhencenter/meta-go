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

var AppPublisherFields = struct {
	ContentID string
	IconURL   string
	ID        string
	Name      string
	Platform  string
	StoreName string
	StoreURL  string
}{
	ContentID: "content_id",
	IconURL:   "icon_url",
	ID:        "id",
	Name:      "name",
	Platform:  "platform",
	StoreName: "store_name",
	StoreURL:  "store_url",
}

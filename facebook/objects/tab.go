package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Tab struct {
	Application               *Application `json:"application,omitempty"`
	CustomImageURL            *string      `json:"custom_image_url,omitempty"`
	CustomName                *string      `json:"custom_name,omitempty"`
	ID                        *core.ID     `json:"id,omitempty"`
	ImageURL                  *string      `json:"image_url,omitempty"`
	IsNonConnectionLandingTab *bool        `json:"is_non_connection_landing_tab,omitempty"`
	IsPermanent               *bool        `json:"is_permanent,omitempty"`
	Link                      *string      `json:"link,omitempty"`
	Name                      *string      `json:"name,omitempty"`
	Position                  *uint64      `json:"position,omitempty"`
}

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

var TabFields = struct {
	Application               string
	CustomImageURL            string
	CustomName                string
	ID                        string
	ImageURL                  string
	IsNonConnectionLandingTab string
	IsPermanent               string
	Link                      string
	Name                      string
	Position                  string
}{
	Application:               "application",
	CustomImageURL:            "custom_image_url",
	CustomName:                "custom_name",
	ID:                        "id",
	ImageURL:                  "image_url",
	IsNonConnectionLandingTab: "is_non_connection_landing_tab",
	IsPermanent:               "is_permanent",
	Link:                      "link",
	Name:                      "name",
	Position:                  "position",
}

package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PartnerAccountLinking struct {
	Adaccount             *AdAccount   `json:"adaccount,omitempty"`
	App                   *Application `json:"app,omitempty"`
	Business              *Business    `json:"business,omitempty"`
	Externalidentifier    *string      `json:"externalidentifier,omitempty"`
	Externalidentifieruri *string      `json:"externalidentifieruri,omitempty"`
	ID                    *core.ID     `json:"id,omitempty"`
	Partnername           *string      `json:"partnername,omitempty"`
	Pixel                 *string      `json:"pixel,omitempty"`
}

var PartnerAccountLinkingFields = struct {
	Adaccount             string
	App                   string
	Business              string
	Externalidentifier    string
	Externalidentifieruri string
	ID                    string
	Partnername           string
	Pixel                 string
}{
	Adaccount:             "adaccount",
	App:                   "app",
	Business:              "business",
	Externalidentifier:    "externalidentifier",
	Externalidentifieruri: "externalidentifieruri",
	ID:                    "id",
	Partnername:           "partnername",
	Pixel:                 "pixel",
}

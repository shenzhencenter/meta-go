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

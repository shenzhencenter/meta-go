package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdColumnSizes struct {
	AdmarketAccount *AdAccount           `json:"admarket_account,omitempty"`
	AppID           *core.ID             `json:"app_id,omitempty"`
	Columns         *[]map[string]string `json:"columns,omitempty"`
	ID              *core.ID             `json:"id,omitempty"`
	Owner           *User                `json:"owner,omitempty"`
	Page            *string              `json:"page,omitempty"`
	Report          *string              `json:"report,omitempty"`
	Tab             *string              `json:"tab,omitempty"`
	View            *string              `json:"view,omitempty"`
}

var AdColumnSizesFields = struct {
	AdmarketAccount string
	AppID           string
	Columns         string
	ID              string
	Owner           string
	Page            string
	Report          string
	Tab             string
	View            string
}{
	AdmarketAccount: "admarket_account",
	AppID:           "app_id",
	Columns:         "columns",
	ID:              "id",
	Owner:           "owner",
	Page:            "page",
	Report:          "report",
	Tab:             "tab",
	View:            "view",
}

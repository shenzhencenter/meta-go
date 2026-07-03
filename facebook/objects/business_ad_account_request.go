package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessAdAccountRequest struct {
	AdAccount *AdAccount `json:"ad_account,omitempty"`
	ID        *core.ID   `json:"id,omitempty"`
}

var BusinessAdAccountRequestFields = struct {
	AdAccount string
	ID        string
}{
	AdAccount: "ad_account",
	ID:        "id",
}

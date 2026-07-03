package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type BusinessAdAccountRequest struct {
	AdAccount *AdAccount `json:"ad_account,omitempty"`
	ID        *core.ID   `json:"id,omitempty"`
}

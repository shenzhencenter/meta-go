package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsTextSuggestions struct {
	AdAccountID          *core.ID                  `json:"ad_account_id,omitempty"`
	Bodies               *[]map[string]interface{} `json:"bodies,omitempty"`
	Descriptions         *[]map[string]interface{} `json:"descriptions,omitempty"`
	InactiveSessionTally *int                      `json:"inactive_session_tally,omitempty"`
	Long                 *[]map[string]interface{} `json:"long,omitempty"`
	Short                *[]map[string]interface{} `json:"short,omitempty"`
	Titles               *[]map[string]interface{} `json:"titles,omitempty"`
}

package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type MediaCopyrightUpdateRecord struct {
	ActionTypes         *[]string                 `json:"action_types,omitempty"`
	Actor               *User                     `json:"actor,omitempty"`
	ActorType           *string                   `json:"actor_type,omitempty"`
	CreationTime        *time.Time                `json:"creation_time,omitempty"`
	ID                  *core.ID                  `json:"id,omitempty"`
	OwnershipCountries  *[]map[string]interface{} `json:"ownership_countries,omitempty"`
	WhitelistedAccounts *[]map[string]interface{} `json:"whitelisted_accounts,omitempty"`
}

package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type MediaCopyrightUpdateRecord struct {
	ActionTypes         *[]string                 `json:"action_types,omitempty"`
	Actor               *User                     `json:"actor,omitempty"`
	ActorType           *string                   `json:"actor_type,omitempty"`
	CreationTime        *core.Time                `json:"creation_time,omitempty"`
	ID                  *core.ID                  `json:"id,omitempty"`
	OwnershipCountries  *[]map[string]interface{} `json:"ownership_countries,omitempty"`
	WhitelistedAccounts *[]map[string]interface{} `json:"whitelisted_accounts,omitempty"`
}

var MediaCopyrightUpdateRecordFields = struct {
	ActionTypes         string
	Actor               string
	ActorType           string
	CreationTime        string
	ID                  string
	OwnershipCountries  string
	WhitelistedAccounts string
}{
	ActionTypes:         "action_types",
	Actor:               "actor",
	ActorType:           "actor_type",
	CreationTime:        "creation_time",
	ID:                  "id",
	OwnershipCountries:  "ownership_countries",
	WhitelistedAccounts: "whitelisted_accounts",
}

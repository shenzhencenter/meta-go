package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsCreationSavedState struct {
	AdAccount           *AdAccount              `json:"ad_account,omitempty"`
	ID                  *core.ID                `json:"id,omitempty"`
	SerializedStoreData *string                 `json:"serialized_store_data,omitempty"`
	TimeUpdated         *core.Time              `json:"time_updated,omitempty"`
	User                *map[string]interface{} `json:"user,omitempty"`
}

var AdsCreationSavedStateFields = struct {
	AdAccount           string
	ID                  string
	SerializedStoreData string
	TimeUpdated         string
	User                string
}{
	AdAccount:           "ad_account",
	ID:                  "id",
	SerializedStoreData: "serialized_store_data",
	TimeUpdated:         "time_updated",
	User:                "user",
}

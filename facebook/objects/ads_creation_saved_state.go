package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AdsCreationSavedState struct {
	AdAccount           *AdAccount              `json:"ad_account,omitempty"`
	ID                  *core.ID                `json:"id,omitempty"`
	SerializedStoreData *string                 `json:"serialized_store_data,omitempty"`
	TimeUpdated         *time.Time              `json:"time_updated,omitempty"`
	User                *map[string]interface{} `json:"user,omitempty"`
}

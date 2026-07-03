package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type CopyrightOwnershipTransfer struct {
	Assets                  *[]map[string]interface{} `json:"assets,omitempty"`
	HasOwnershipBeenUpdated *bool                     `json:"has_ownership_been_updated,omitempty"`
	ID                      *core.ID                  `json:"id,omitempty"`
	NumAssets               *int                      `json:"num_assets,omitempty"`
	ReceivingRightsHolder   *Profile                  `json:"receiving_rights_holder,omitempty"`
	SendingRightsHolder     *Profile                  `json:"sending_rights_holder,omitempty"`
	Status                  *string                   `json:"status,omitempty"`
	TransferTerritories     *[]string                 `json:"transfer_territories,omitempty"`
	TransferTime            *time.Time                `json:"transfer_time,omitempty"`
}

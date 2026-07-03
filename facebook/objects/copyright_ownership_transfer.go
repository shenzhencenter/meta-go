package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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
	TransferTime            *core.Time                `json:"transfer_time,omitempty"`
}

var CopyrightOwnershipTransferFields = struct {
	Assets                  string
	HasOwnershipBeenUpdated string
	ID                      string
	NumAssets               string
	ReceivingRightsHolder   string
	SendingRightsHolder     string
	Status                  string
	TransferTerritories     string
	TransferTime            string
}{
	Assets:                  "assets",
	HasOwnershipBeenUpdated: "has_ownership_been_updated",
	ID:                      "id",
	NumAssets:               "num_assets",
	ReceivingRightsHolder:   "receiving_rights_holder",
	SendingRightsHolder:     "sending_rights_holder",
	Status:                  "status",
	TransferTerritories:     "transfer_territories",
	TransferTime:            "transfer_time",
}

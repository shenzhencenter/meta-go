package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WifiInformation struct {
	ID                *core.ID `json:"id,omitempty"`
	Name              *string  `json:"name,omitempty"`
	NetworkAccessType *string  `json:"network_access_type,omitempty"`
}

var WifiInformationFields = struct {
	ID                string
	Name              string
	NetworkAccessType string
}{
	ID:                "id",
	Name:              "name",
	NetworkAccessType: "network_access_type",
}

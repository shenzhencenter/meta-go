package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessAssetSharingAgreement struct {
	ID               *core.ID  `json:"id,omitempty"`
	Initiator        *Business `json:"initiator,omitempty"`
	Recipient        *Business `json:"recipient,omitempty"`
	RelationshipType *[]string `json:"relationship_type,omitempty"`
	RequestStatus    *string   `json:"request_status,omitempty"`
	RequestType      *string   `json:"request_type,omitempty"`
}

var BusinessAssetSharingAgreementFields = struct {
	ID               string
	Initiator        string
	Recipient        string
	RelationshipType string
	RequestStatus    string
	RequestType      string
}{
	ID:               "id",
	Initiator:        "initiator",
	Recipient:        "recipient",
	RelationshipType: "relationship_type",
	RequestStatus:    "request_status",
	RequestType:      "request_type",
}

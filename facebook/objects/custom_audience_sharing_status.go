package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CustomAudienceSharingStatus struct {
	SharingRelationshipID *core.ID `json:"sharing_relationship_id,omitempty"`
	Status                *string  `json:"status,omitempty"`
}

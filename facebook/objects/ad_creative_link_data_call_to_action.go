package objects

import (
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type AdCreativeLinkDataCallToAction struct {
	Type  *enums.AdcreativelinkdatacalltoactionType `json:"type,omitempty"`
	Value *AdCreativeLinkDataCallToActionValue      `json:"value,omitempty"`
}

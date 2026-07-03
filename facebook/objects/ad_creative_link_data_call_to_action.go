package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdCreativeLinkDataCallToAction struct {
	Type  *enums.AdcreativelinkdatacalltoactionType `json:"type,omitempty"`
	Value *AdCreativeLinkDataCallToActionValue      `json:"value,omitempty"`
}

var AdCreativeLinkDataCallToActionFields = struct {
	Type  string
	Value string
}{
	Type:  "type",
	Value: "value",
}

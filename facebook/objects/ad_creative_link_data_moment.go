package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type AdCreativeLinkDataMoment struct {
	ID   *core.ID                            `json:"id,omitempty"`
	Type *enums.AdcreativelinkdatamomentType `json:"type,omitempty"`
}

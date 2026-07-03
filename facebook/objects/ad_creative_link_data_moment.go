package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdCreativeLinkDataMoment struct {
	ID   *core.ID                            `json:"id,omitempty"`
	Type *enums.AdcreativelinkdatamomentType `json:"type,omitempty"`
}

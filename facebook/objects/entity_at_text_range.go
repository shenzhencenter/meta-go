package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type EntityAtTextRange struct {
	ID     *core.ID                     `json:"id,omitempty"`
	Length *uint64                      `json:"length,omitempty"`
	Name   *string                      `json:"name,omitempty"`
	Object *Profile                     `json:"object,omitempty"`
	Offset *uint64                      `json:"offset,omitempty"`
	Type   *enums.EntityattextrangeType `json:"type,omitempty"`
}

var EntityAtTextRangeFields = struct {
	ID     string
	Length string
	Name   string
	Object string
	Offset string
	Type   string
}{
	ID:     "id",
	Length: "length",
	Name:   "name",
	Object: "object",
	Offset: "offset",
	Type:   "type",
}

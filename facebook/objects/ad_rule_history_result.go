package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdRuleHistoryResult struct {
	Actions    *[]AdRuleHistoryResultAction         `json:"actions,omitempty"`
	ObjectID   *core.ID                             `json:"object_id,omitempty"`
	ObjectType *enums.AdrulehistoryresultObjectType `json:"object_type,omitempty"`
}

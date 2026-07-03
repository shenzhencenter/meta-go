package objects

import (
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type AdAccountDefaultObjective struct {
	DefaultObjectiveForUser *enums.AdaccountdefaultobjectiveDefaultObjectiveForUser `json:"default_objective_for_user,omitempty"`
	ObjectiveForLevel       *enums.AdaccountdefaultobjectiveObjectiveForLevel       `json:"objective_for_level,omitempty"`
}

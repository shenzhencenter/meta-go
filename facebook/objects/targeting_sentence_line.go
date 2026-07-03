package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type TargetingSentenceLine struct {
	ID                     *core.ID       `json:"id,omitempty"`
	Params                 *Targeting     `json:"params,omitempty"`
	Targetingsentencelines *[]interface{} `json:"targetingsentencelines,omitempty"`
}

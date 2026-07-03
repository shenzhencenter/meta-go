package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type TargetingSentenceLine struct {
	ID                     *core.ID       `json:"id,omitempty"`
	Params                 *Targeting     `json:"params,omitempty"`
	Targetingsentencelines *[]interface{} `json:"targetingsentencelines,omitempty"`
}

var TargetingSentenceLineFields = struct {
	ID                     string
	Params                 string
	Targetingsentencelines string
}{
	ID:                     "id",
	Params:                 "params",
	Targetingsentencelines: "targetingsentencelines",
}

package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type RevSharePolicy struct {
	PolicyID   *core.ID `json:"policy_id,omitempty"`
	PolicyName *string  `json:"policy_name,omitempty"`
}

var RevSharePolicyFields = struct {
	PolicyID   string
	PolicyName string
}{
	PolicyID:   "policy_id",
	PolicyName: "policy_name",
}

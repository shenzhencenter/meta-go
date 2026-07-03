package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type RevSharePolicy struct {
	PolicyID   *core.ID `json:"policy_id,omitempty"`
	PolicyName *string  `json:"policy_name,omitempty"`
}

package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type RevSharePolicy struct {
	PolicyID   *core.ID `json:"policy_id,omitempty"`
	PolicyName *string  `json:"policy_name,omitempty"`
}

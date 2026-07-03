package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ResearchPollStudy struct {
	Account *AdAccount `json:"account,omitempty"`
	ID      *core.ID   `json:"id,omitempty"`
	Name    *string    `json:"name,omitempty"`
}

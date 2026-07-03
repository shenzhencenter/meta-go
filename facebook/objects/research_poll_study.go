package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ResearchPollStudy struct {
	Account *AdAccount `json:"account,omitempty"`
	ID      *core.ID   `json:"id,omitempty"`
	Name    *string    `json:"name,omitempty"`
}

var ResearchPollStudyFields = struct {
	Account string
	ID      string
	Name    string
}{
	Account: "account",
	ID:      "id",
	Name:    "name",
}

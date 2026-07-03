package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdSavedKeywords struct {
	Account     *AdAccount  `json:"account,omitempty"`
	ID          *core.ID    `json:"id,omitempty"`
	Keywords    *AdKeywords `json:"keywords,omitempty"`
	Name        *string     `json:"name,omitempty"`
	RunStatus   *string     `json:"run_status,omitempty"`
	TimeCreated *core.Time  `json:"time_created,omitempty"`
	TimeUpdated *core.Time  `json:"time_updated,omitempty"`
}

var AdSavedKeywordsFields = struct {
	Account     string
	ID          string
	Keywords    string
	Name        string
	RunStatus   string
	TimeCreated string
	TimeUpdated string
}{
	Account:     "account",
	ID:          "id",
	Keywords:    "keywords",
	Name:        "name",
	RunStatus:   "run_status",
	TimeCreated: "time_created",
	TimeUpdated: "time_updated",
}

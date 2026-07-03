package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type AdSavedKeywords struct {
	Account     *AdAccount  `json:"account,omitempty"`
	ID          *core.ID    `json:"id,omitempty"`
	Keywords    *AdKeywords `json:"keywords,omitempty"`
	Name        *string     `json:"name,omitempty"`
	RunStatus   *string     `json:"run_status,omitempty"`
	TimeCreated *time.Time  `json:"time_created,omitempty"`
	TimeUpdated *time.Time  `json:"time_updated,omitempty"`
}

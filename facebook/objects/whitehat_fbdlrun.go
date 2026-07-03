package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type WhitehatFBDLRun struct {
	CreationTime *time.Time           `json:"creation_time,omitempty"`
	ID           *core.ID             `json:"id,omitempty"`
	IsPinned     *bool                `json:"is_pinned,omitempty"`
	Note         *string              `json:"note,omitempty"`
	Result       *[]map[string]string `json:"result,omitempty"`
	RunCode      *string              `json:"run_code,omitempty"`
	Status       *string              `json:"status,omitempty"`
	UserType     *string              `json:"user_type,omitempty"`
}

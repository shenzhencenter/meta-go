package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WhitehatFBDLRun struct {
	CreationTime *core.Time           `json:"creation_time,omitempty"`
	ID           *core.ID             `json:"id,omitempty"`
	IsPinned     *bool                `json:"is_pinned,omitempty"`
	Note         *string              `json:"note,omitempty"`
	Result       *[]map[string]string `json:"result,omitempty"`
	RunCode      *string              `json:"run_code,omitempty"`
	Status       *string              `json:"status,omitempty"`
	UserType     *string              `json:"user_type,omitempty"`
}

var WhitehatFBDLRunFields = struct {
	CreationTime string
	ID           string
	IsPinned     string
	Note         string
	Result       string
	RunCode      string
	Status       string
	UserType     string
}{
	CreationTime: "creation_time",
	ID:           "id",
	IsPinned:     "is_pinned",
	Note:         "note",
	Result:       "result",
	RunCode:      "run_code",
	Status:       "status",
	UserType:     "user_type",
}

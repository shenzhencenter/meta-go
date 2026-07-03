package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdgroupFacebookFeedback struct {
	ID      *core.ID `json:"id,omitempty"`
	Preview *string  `json:"preview,omitempty"`
}

var AdgroupFacebookFeedbackFields = struct {
	ID      string
	Preview string
}{
	ID:      "id",
	Preview: "preview",
}

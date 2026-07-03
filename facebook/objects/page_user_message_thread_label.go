package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type PageUserMessageThreadLabel struct {
	ID            *core.ID `json:"id,omitempty"`
	PageLabelName *string  `json:"page_label_name,omitempty"`
}

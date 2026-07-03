package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PageUserMessageThreadLabel struct {
	ID            *core.ID `json:"id,omitempty"`
	PageLabelName *string  `json:"page_label_name,omitempty"`
}

var PageUserMessageThreadLabelFields = struct {
	ID            string
	PageLabelName string
}{
	ID:            "id",
	PageLabelName: "page_label_name",
}

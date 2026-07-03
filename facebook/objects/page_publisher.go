package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PagePublisher struct {
	GlobalParentID *core.ID `json:"global_parent_id,omitempty"`
	Icon           *string  `json:"icon,omitempty"`
	ID             *core.ID `json:"id,omitempty"`
	Name           *string  `json:"name,omitempty"`
	URL            *string  `json:"url,omitempty"`
}

var PagePublisherFields = struct {
	GlobalParentID string
	Icon           string
	ID             string
	Name           string
	URL            string
}{
	GlobalParentID: "global_parent_id",
	Icon:           "icon",
	ID:             "id",
	Name:           "name",
	URL:            "url",
}

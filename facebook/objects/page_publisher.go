package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type PagePublisher struct {
	GlobalParentID *core.ID `json:"global_parent_id,omitempty"`
	Icon           *string  `json:"icon,omitempty"`
	ID             *core.ID `json:"id,omitempty"`
	Name           *string  `json:"name,omitempty"`
	URL            *string  `json:"url,omitempty"`
}

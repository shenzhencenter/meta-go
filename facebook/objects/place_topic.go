package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PlaceTopic struct {
	Count            *uint64   `json:"count,omitempty"`
	HasChildren      *bool     `json:"has_children,omitempty"`
	IconURL          *string   `json:"icon_url,omitempty"`
	ID               *core.ID  `json:"id,omitempty"`
	Name             *string   `json:"name,omitempty"`
	PluralName       *string   `json:"plural_name,omitempty"`
	TopSubtopicNames *[]string `json:"top_subtopic_names,omitempty"`
}

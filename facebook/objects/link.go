package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Link struct {
	Caption             *string                 `json:"caption,omitempty"`
	CreatedTime         *core.Time              `json:"created_time,omitempty"`
	Description         *string                 `json:"description,omitempty"`
	From                *map[string]interface{} `json:"from,omitempty"`
	Icon                *string                 `json:"icon,omitempty"`
	ID                  *core.ID                `json:"id,omitempty"`
	Link                *string                 `json:"link,omitempty"`
	Message             *string                 `json:"message,omitempty"`
	MultiShareOptimized *bool                   `json:"multi_share_optimized,omitempty"`
	Name                *string                 `json:"name,omitempty"`
	Privacy             *Privacy                `json:"privacy,omitempty"`
	Via                 *map[string]interface{} `json:"via,omitempty"`
}

var LinkFields = struct {
	Caption             string
	CreatedTime         string
	Description         string
	From                string
	Icon                string
	ID                  string
	Link                string
	Message             string
	MultiShareOptimized string
	Name                string
	Privacy             string
	Via                 string
}{
	Caption:             "caption",
	CreatedTime:         "created_time",
	Description:         "description",
	From:                "from",
	Icon:                "icon",
	ID:                  "id",
	Link:                "link",
	Message:             "message",
	MultiShareOptimized: "multi_share_optimized",
	Name:                "name",
	Privacy:             "privacy",
	Via:                 "via",
}

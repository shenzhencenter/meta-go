package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type Link struct {
	Caption             *string                 `json:"caption,omitempty"`
	CreatedTime         *time.Time              `json:"created_time,omitempty"`
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

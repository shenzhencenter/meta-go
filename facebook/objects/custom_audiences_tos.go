package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CustomAudiencesTOS struct {
	Content *string  `json:"content,omitempty"`
	ID      *core.ID `json:"id,omitempty"`
	Type    *string  `json:"type,omitempty"`
}

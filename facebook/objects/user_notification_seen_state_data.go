package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type UserNotificationSeenStateData struct {
	ID        *core.ID `json:"id,omitempty"`
	SeenState *string  `json:"seen_state,omitempty"`
}

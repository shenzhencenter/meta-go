package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type UserNotificationSeenStateData struct {
	ID        *core.ID `json:"id,omitempty"`
	SeenState *string  `json:"seen_state,omitempty"`
}

var UserNotificationSeenStateDataFields = struct {
	ID        string
	SeenState string
}{
	ID:        "id",
	SeenState: "seen_state",
}

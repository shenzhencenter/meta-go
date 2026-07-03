package objects

import (
	"time"
)

type BlockedUserInfo struct {
	BlockTime *time.Time `json:"block_time,omitempty"`
	BlockType *string    `json:"block_type,omitempty"`
	Fbid      *string    `json:"fbid,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Username  *string    `json:"username,omitempty"`
}

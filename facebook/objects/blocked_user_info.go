package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BlockedUserInfo struct {
	BlockTime *core.Time `json:"block_time,omitempty"`
	BlockType *string    `json:"block_type,omitempty"`
	Fbid      *string    `json:"fbid,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Username  *string    `json:"username,omitempty"`
}

var BlockedUserInfoFields = struct {
	BlockTime string
	BlockType string
	Fbid      string
	Name      string
	Username  string
}{
	BlockTime: "block_time",
	BlockType: "block_type",
	Fbid:      "fbid",
	Name:      "name",
	Username:  "username",
}

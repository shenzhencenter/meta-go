package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ThreadsUser struct {
	ThreadsUserID         *core.ID `json:"threads_user_id,omitempty"`
	ThreadsUserProfilePic *string  `json:"threads_user_profile_pic,omitempty"`
}

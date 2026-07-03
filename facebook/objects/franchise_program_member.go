package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type FranchiseProgramMember struct {
	Business         *Business  `json:"business,omitempty"`
	EndDate          *time.Time `json:"end_date,omitempty"`
	ID               *core.ID   `json:"id,omitempty"`
	JoinDate         *time.Time `json:"join_date,omitempty"`
	MemberAdAccount  *AdAccount `json:"member_ad_account,omitempty"`
	MemberUser       *User      `json:"member_user,omitempty"`
	MembershipStatus *string    `json:"membership_status,omitempty"`
	Page             *Page      `json:"page,omitempty"`
}

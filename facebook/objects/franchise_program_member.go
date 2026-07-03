package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type FranchiseProgramMember struct {
	Business         *Business  `json:"business,omitempty"`
	EndDate          *core.Time `json:"end_date,omitempty"`
	ID               *core.ID   `json:"id,omitempty"`
	JoinDate         *core.Time `json:"join_date,omitempty"`
	MemberAdAccount  *AdAccount `json:"member_ad_account,omitempty"`
	MemberUser       *User      `json:"member_user,omitempty"`
	MembershipStatus *string    `json:"membership_status,omitempty"`
	Page             *Page      `json:"page,omitempty"`
}

var FranchiseProgramMemberFields = struct {
	Business         string
	EndDate          string
	ID               string
	JoinDate         string
	MemberAdAccount  string
	MemberUser       string
	MembershipStatus string
	Page             string
}{
	Business:         "business",
	EndDate:          "end_date",
	ID:               "id",
	JoinDate:         "join_date",
	MemberAdAccount:  "member_ad_account",
	MemberUser:       "member_user",
	MembershipStatus: "membership_status",
	Page:             "page",
}

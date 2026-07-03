package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type VideoCopyrightMatch struct {
	CreatedDate             *time.Time                `json:"created_date,omitempty"`
	ID                      *core.ID                  `json:"id,omitempty"`
	LastModifiedUser        *User                     `json:"last_modified_user,omitempty"`
	MatchData               *[]map[string]interface{} `json:"match_data,omitempty"`
	MatchStatus             *string                   `json:"match_status,omitempty"`
	Notes                   *string                   `json:"notes,omitempty"`
	Permalink               *string                   `json:"permalink,omitempty"`
	PolicyEvalModifyReasons *[]map[string]interface{} `json:"policy_eval_modify_reasons,omitempty"`
	UgcContentFormat        *string                   `json:"ugc_content_format,omitempty"`
}

package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type VideoCopyrightMatch struct {
	CreatedDate             *core.Time                `json:"created_date,omitempty"`
	ID                      *core.ID                  `json:"id,omitempty"`
	LastModifiedUser        *User                     `json:"last_modified_user,omitempty"`
	MatchData               *[]map[string]interface{} `json:"match_data,omitempty"`
	MatchStatus             *string                   `json:"match_status,omitempty"`
	Notes                   *string                   `json:"notes,omitempty"`
	Permalink               *string                   `json:"permalink,omitempty"`
	PolicyEvalModifyReasons *[]map[string]interface{} `json:"policy_eval_modify_reasons,omitempty"`
	UgcContentFormat        *string                   `json:"ugc_content_format,omitempty"`
}

var VideoCopyrightMatchFields = struct {
	CreatedDate             string
	ID                      string
	LastModifiedUser        string
	MatchData               string
	MatchStatus             string
	Notes                   string
	Permalink               string
	PolicyEvalModifyReasons string
	UgcContentFormat        string
}{
	CreatedDate:             "created_date",
	ID:                      "id",
	LastModifiedUser:        "last_modified_user",
	MatchData:               "match_data",
	MatchStatus:             "match_status",
	Notes:                   "notes",
	Permalink:               "permalink",
	PolicyEvalModifyReasons: "policy_eval_modify_reasons",
	UgcContentFormat:        "ugc_content_format",
}

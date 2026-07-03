package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type MusicWorkCopyright struct {
	AvailableUiActions                        *[]string           `json:"available_ui_actions,omitempty"`
	ClaimStatus                               *string             `json:"claim_status,omitempty"`
	CreationTime                              *core.Time          `json:"creation_time,omitempty"`
	DisplayedFbMatchesCount                   *int                `json:"displayed_fb_matches_count,omitempty"`
	DisplayedIgMatchesCount                   *int                `json:"displayed_ig_matches_count,omitempty"`
	DisplayedMatchesCount                     *int                `json:"displayed_matches_count,omitempty"`
	HasRevShareEligibleIsrcs                  *bool               `json:"has_rev_share_eligible_isrcs,omitempty"`
	ID                                        *core.ID            `json:"id,omitempty"`
	IsLinkingRequiredToMonetizeForManualClaim *bool               `json:"is_linking_required_to_monetize_for_manual_claim,omitempty"`
	MatchRule                                 *VideoCopyrightRule `json:"match_rule,omitempty"`
	Status                                    *string             `json:"status,omitempty"`
	Tags                                      *[]string           `json:"tags,omitempty"`
	UpdateTime                                *core.Time          `json:"update_time,omitempty"`
}

var MusicWorkCopyrightFields = struct {
	AvailableUiActions                        string
	ClaimStatus                               string
	CreationTime                              string
	DisplayedFbMatchesCount                   string
	DisplayedIgMatchesCount                   string
	DisplayedMatchesCount                     string
	HasRevShareEligibleIsrcs                  string
	ID                                        string
	IsLinkingRequiredToMonetizeForManualClaim string
	MatchRule                                 string
	Status                                    string
	Tags                                      string
	UpdateTime                                string
}{
	AvailableUiActions:       "available_ui_actions",
	ClaimStatus:              "claim_status",
	CreationTime:             "creation_time",
	DisplayedFbMatchesCount:  "displayed_fb_matches_count",
	DisplayedIgMatchesCount:  "displayed_ig_matches_count",
	DisplayedMatchesCount:    "displayed_matches_count",
	HasRevShareEligibleIsrcs: "has_rev_share_eligible_isrcs",
	ID:                       "id",
	IsLinkingRequiredToMonetizeForManualClaim: "is_linking_required_to_monetize_for_manual_claim",
	MatchRule:  "match_rule",
	Status:     "status",
	Tags:       "tags",
	UpdateTime: "update_time",
}

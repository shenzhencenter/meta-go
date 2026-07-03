package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type MusicWorkCopyright struct {
	AvailableUiActions                        *[]string           `json:"available_ui_actions,omitempty"`
	ClaimStatus                               *string             `json:"claim_status,omitempty"`
	CreationTime                              *time.Time          `json:"creation_time,omitempty"`
	DisplayedFbMatchesCount                   *int                `json:"displayed_fb_matches_count,omitempty"`
	DisplayedIgMatchesCount                   *int                `json:"displayed_ig_matches_count,omitempty"`
	DisplayedMatchesCount                     *int                `json:"displayed_matches_count,omitempty"`
	HasRevShareEligibleIsrcs                  *bool               `json:"has_rev_share_eligible_isrcs,omitempty"`
	ID                                        *core.ID            `json:"id,omitempty"`
	IsLinkingRequiredToMonetizeForManualClaim *bool               `json:"is_linking_required_to_monetize_for_manual_claim,omitempty"`
	MatchRule                                 *VideoCopyrightRule `json:"match_rule,omitempty"`
	Status                                    *string             `json:"status,omitempty"`
	Tags                                      *[]string           `json:"tags,omitempty"`
	UpdateTime                                *time.Time          `json:"update_time,omitempty"`
}

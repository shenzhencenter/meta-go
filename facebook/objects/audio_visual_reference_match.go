package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AudioVisualReferenceMatch struct {
	AudioConflictingSegments                 *[]map[string]interface{}              `json:"audio_conflicting_segments,omitempty"`
	AudioCurrentConflictResolvedSegments     *[]map[string]interface{}              `json:"audio_current_conflict_resolved_segments,omitempty"`
	AudioSegmentResolutionHistory            *[]map[string]interface{}              `json:"audio_segment_resolution_history,omitempty"`
	ConflictStatus                           *string                                `json:"conflict_status,omitempty"`
	ConflictType                             *string                                `json:"conflict_type,omitempty"`
	ConflictingCountries                     *[]string                              `json:"conflicting_countries,omitempty"`
	CountryResolutionHistory                 *[]map[string][]map[string]interface{} `json:"country_resolution_history,omitempty"`
	CreationTime                             *time.Time                             `json:"creation_time,omitempty"`
	CurrentConflictResolvedCountries         *[]map[string]map[string]interface{}   `json:"current_conflict_resolved_countries,omitempty"`
	DisplayedMatchState                      *string                                `json:"displayed_match_state,omitempty"`
	DisputeFormDataEntriesWithTranslations   *[]map[string]interface{}              `json:"dispute_form_data_entries_with_translations,omitempty"`
	ExpirationTime                           *time.Time                             `json:"expiration_time,omitempty"`
	ID                                       *core.ID                               `json:"id,omitempty"`
	IsDisputable                             *bool                                  `json:"is_disputable,omitempty"`
	MatchState                               *string                                `json:"match_state,omitempty"`
	MatchedOverlapPercentage                 *float64                               `json:"matched_overlap_percentage,omitempty"`
	MatchedOwnerMatchDurationInSec           *float64                               `json:"matched_owner_match_duration_in_sec,omitempty"`
	MatchedReferenceOwner                    *Profile                               `json:"matched_reference_owner,omitempty"`
	ModificationHistory                      *[]map[string]interface{}              `json:"modification_history,omitempty"`
	NumMatchesOnMatchedSide                  *uint64                                `json:"num_matches_on_matched_side,omitempty"`
	NumMatchesOnRefSide                      *uint64                                `json:"num_matches_on_ref_side,omitempty"`
	RefOwnerMatchDurationInSec               *float64                               `json:"ref_owner_match_duration_in_sec,omitempty"`
	ReferenceOverlapPercentage               *float64                               `json:"reference_overlap_percentage,omitempty"`
	ReferenceOwner                           *Profile                               `json:"reference_owner,omitempty"`
	RejectionFormDataEntriesWithTranslations *[]map[string]interface{}              `json:"rejection_form_data_entries_with_translations,omitempty"`
	ResolutionDetails                        *string                                `json:"resolution_details,omitempty"`
	ResolutionReason                         *string                                `json:"resolution_reason,omitempty"`
	UpdateTime                               *time.Time                             `json:"update_time,omitempty"`
	ViewsOnMatchedSide                       *uint64                                `json:"views_on_matched_side,omitempty"`
	VisualConflictingSegments                *[]map[string]interface{}              `json:"visual_conflicting_segments,omitempty"`
	VisualCurrentConflictResolvedSegments    *[]map[string]interface{}              `json:"visual_current_conflict_resolved_segments,omitempty"`
	VisualSegmentResolutionHistory           *[]map[string]interface{}              `json:"visual_segment_resolution_history,omitempty"`
}

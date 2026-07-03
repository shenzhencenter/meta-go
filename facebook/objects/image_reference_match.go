package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type ImageReferenceMatch struct {
	ConflictStatus                           *string                                `json:"conflict_status,omitempty"`
	ConflictingCountries                     *[]string                              `json:"conflicting_countries,omitempty"`
	CountryResolutionHistory                 *[]map[string][]map[string]interface{} `json:"country_resolution_history,omitempty"`
	CreationTime                             *time.Time                             `json:"creation_time,omitempty"`
	CurrentConflictResolvedCountries         *[]map[string]map[string]interface{}   `json:"current_conflict_resolved_countries,omitempty"`
	DisplayedMatchState                      *string                                `json:"displayed_match_state,omitempty"`
	DisputeFormDataEntriesWithTranslations   *[]map[string]interface{}              `json:"dispute_form_data_entries_with_translations,omitempty"`
	ExpirationTime                           *time.Time                             `json:"expiration_time,omitempty"`
	ID                                       *core.ID                               `json:"id,omitempty"`
	MatchState                               *string                                `json:"match_state,omitempty"`
	MatchedReferenceCopyright                *ImageCopyright                        `json:"matched_reference_copyright,omitempty"`
	MatchedReferenceOwner                    *Profile                               `json:"matched_reference_owner,omitempty"`
	ModificationHistory                      *[]map[string]interface{}              `json:"modification_history,omitempty"`
	ReferenceCopyright                       *ImageCopyright                        `json:"reference_copyright,omitempty"`
	ReferenceOwner                           *Profile                               `json:"reference_owner,omitempty"`
	RejectionFormDataEntriesWithTranslations *[]map[string]interface{}              `json:"rejection_form_data_entries_with_translations,omitempty"`
	ResolutionReason                         *string                                `json:"resolution_reason,omitempty"`
	UpdateTime                               *time.Time                             `json:"update_time,omitempty"`
}

package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ImageReferenceMatch struct {
	ConflictStatus                           *string                                `json:"conflict_status,omitempty"`
	ConflictingCountries                     *[]string                              `json:"conflicting_countries,omitempty"`
	CountryResolutionHistory                 *[]map[string][]map[string]interface{} `json:"country_resolution_history,omitempty"`
	CreationTime                             *core.Time                             `json:"creation_time,omitempty"`
	CurrentConflictResolvedCountries         *[]map[string]map[string]interface{}   `json:"current_conflict_resolved_countries,omitempty"`
	DisplayedMatchState                      *string                                `json:"displayed_match_state,omitempty"`
	DisputeFormDataEntriesWithTranslations   *[]map[string]interface{}              `json:"dispute_form_data_entries_with_translations,omitempty"`
	ExpirationTime                           *core.Time                             `json:"expiration_time,omitempty"`
	ID                                       *core.ID                               `json:"id,omitempty"`
	MatchState                               *string                                `json:"match_state,omitempty"`
	MatchedReferenceCopyright                *ImageCopyright                        `json:"matched_reference_copyright,omitempty"`
	MatchedReferenceOwner                    *Profile                               `json:"matched_reference_owner,omitempty"`
	ModificationHistory                      *[]map[string]interface{}              `json:"modification_history,omitempty"`
	ReferenceCopyright                       *ImageCopyright                        `json:"reference_copyright,omitempty"`
	ReferenceOwner                           *Profile                               `json:"reference_owner,omitempty"`
	RejectionFormDataEntriesWithTranslations *[]map[string]interface{}              `json:"rejection_form_data_entries_with_translations,omitempty"`
	ResolutionReason                         *string                                `json:"resolution_reason,omitempty"`
	UpdateTime                               *core.Time                             `json:"update_time,omitempty"`
}

var ImageReferenceMatchFields = struct {
	ConflictStatus                           string
	ConflictingCountries                     string
	CountryResolutionHistory                 string
	CreationTime                             string
	CurrentConflictResolvedCountries         string
	DisplayedMatchState                      string
	DisputeFormDataEntriesWithTranslations   string
	ExpirationTime                           string
	ID                                       string
	MatchState                               string
	MatchedReferenceCopyright                string
	MatchedReferenceOwner                    string
	ModificationHistory                      string
	ReferenceCopyright                       string
	ReferenceOwner                           string
	RejectionFormDataEntriesWithTranslations string
	ResolutionReason                         string
	UpdateTime                               string
}{
	ConflictStatus:                           "conflict_status",
	ConflictingCountries:                     "conflicting_countries",
	CountryResolutionHistory:                 "country_resolution_history",
	CreationTime:                             "creation_time",
	CurrentConflictResolvedCountries:         "current_conflict_resolved_countries",
	DisplayedMatchState:                      "displayed_match_state",
	DisputeFormDataEntriesWithTranslations:   "dispute_form_data_entries_with_translations",
	ExpirationTime:                           "expiration_time",
	ID:                                       "id",
	MatchState:                               "match_state",
	MatchedReferenceCopyright:                "matched_reference_copyright",
	MatchedReferenceOwner:                    "matched_reference_owner",
	ModificationHistory:                      "modification_history",
	ReferenceCopyright:                       "reference_copyright",
	ReferenceOwner:                           "reference_owner",
	RejectionFormDataEntriesWithTranslations: "rejection_form_data_entries_with_translations",
	ResolutionReason:                         "resolution_reason",
	UpdateTime:                               "update_time",
}

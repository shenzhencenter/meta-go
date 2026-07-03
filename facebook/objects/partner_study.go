package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PartnerStudy struct {
	AdditionalInfo                 *string    `json:"additional_info,omitempty"`
	Brand                          *string    `json:"brand,omitempty"`
	ClientName                     *string    `json:"client_name,omitempty"`
	Emails                         *string    `json:"emails,omitempty"`
	ID                             *core.ID   `json:"id,omitempty"`
	InputIds                       *[]core.ID `json:"input_ids,omitempty"`
	IsExport                       *bool      `json:"is_export,omitempty"`
	LiftStudy                      *AdStudy   `json:"lift_study,omitempty"`
	Location                       *string    `json:"location,omitempty"`
	MatchFileDs                    *string    `json:"match_file_ds,omitempty"`
	Name                           *string    `json:"name,omitempty"`
	PartnerDefinedID               *core.ID   `json:"partner_defined_id,omitempty"`
	PartnerHouseholdGraphDatasetID *core.ID   `json:"partner_household_graph_dataset_id,omitempty"`
	Status                         *string    `json:"status,omitempty"`
	StudyEndDate                   *core.Time `json:"study_end_date,omitempty"`
	StudyStartDate                 *core.Time `json:"study_start_date,omitempty"`
	StudyType                      *string    `json:"study_type,omitempty"`
	SubmitDate                     *core.Time `json:"submit_date,omitempty"`
}

var PartnerStudyFields = struct {
	AdditionalInfo                 string
	Brand                          string
	ClientName                     string
	Emails                         string
	ID                             string
	InputIds                       string
	IsExport                       string
	LiftStudy                      string
	Location                       string
	MatchFileDs                    string
	Name                           string
	PartnerDefinedID               string
	PartnerHouseholdGraphDatasetID string
	Status                         string
	StudyEndDate                   string
	StudyStartDate                 string
	StudyType                      string
	SubmitDate                     string
}{
	AdditionalInfo:                 "additional_info",
	Brand:                          "brand",
	ClientName:                     "client_name",
	Emails:                         "emails",
	ID:                             "id",
	InputIds:                       "input_ids",
	IsExport:                       "is_export",
	LiftStudy:                      "lift_study",
	Location:                       "location",
	MatchFileDs:                    "match_file_ds",
	Name:                           "name",
	PartnerDefinedID:               "partner_defined_id",
	PartnerHouseholdGraphDatasetID: "partner_household_graph_dataset_id",
	Status:                         "status",
	StudyEndDate:                   "study_end_date",
	StudyStartDate:                 "study_start_date",
	StudyType:                      "study_type",
	SubmitDate:                     "submit_date",
}

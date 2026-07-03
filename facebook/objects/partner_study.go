package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
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
	StudyEndDate                   *time.Time `json:"study_end_date,omitempty"`
	StudyStartDate                 *time.Time `json:"study_start_date,omitempty"`
	StudyType                      *string    `json:"study_type,omitempty"`
	SubmitDate                     *time.Time `json:"submit_date,omitempty"`
}

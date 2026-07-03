package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type ThirdPartyPartnerLiftRequest struct {
	AdEntities                     *[]string                                 `json:"ad_entities,omitempty"`
	Country                        *string                                   `json:"country,omitempty"`
	CreatedTime                    *core.Time                                `json:"created_time,omitempty"`
	Description                    *string                                   `json:"description,omitempty"`
	HoldoutSize                    *float64                                  `json:"holdout_size,omitempty"`
	ID                             *core.ID                                  `json:"id,omitempty"`
	LegacyAdsDataPartnerID         *core.ID                                  `json:"legacy_ads_data_partner_id,omitempty"`
	LegacyAdsDataPartnerName       *string                                   `json:"legacy_ads_data_partner_name,omitempty"`
	ModifiedTime                   *core.Time                                `json:"modified_time,omitempty"`
	OwnerInstanceID                *core.ID                                  `json:"owner_instance_id,omitempty"`
	PartnerHouseholdGraphDatasetID *core.ID                                  `json:"partner_household_graph_dataset_id,omitempty"`
	Region                         *string                                   `json:"region,omitempty"`
	Status                         *enums.ThirdpartypartnerliftrequestStatus `json:"status,omitempty"`
	StudyCells                     *[]string                                 `json:"study_cells,omitempty"`
	StudyEndTime                   *core.Time                                `json:"study_end_time,omitempty"`
	StudyStartTime                 *core.Time                                `json:"study_start_time,omitempty"`
}

var ThirdPartyPartnerLiftRequestFields = struct {
	AdEntities                     string
	Country                        string
	CreatedTime                    string
	Description                    string
	HoldoutSize                    string
	ID                             string
	LegacyAdsDataPartnerID         string
	LegacyAdsDataPartnerName       string
	ModifiedTime                   string
	OwnerInstanceID                string
	PartnerHouseholdGraphDatasetID string
	Region                         string
	Status                         string
	StudyCells                     string
	StudyEndTime                   string
	StudyStartTime                 string
}{
	AdEntities:                     "ad_entities",
	Country:                        "country",
	CreatedTime:                    "created_time",
	Description:                    "description",
	HoldoutSize:                    "holdout_size",
	ID:                             "id",
	LegacyAdsDataPartnerID:         "legacy_ads_data_partner_id",
	LegacyAdsDataPartnerName:       "legacy_ads_data_partner_name",
	ModifiedTime:                   "modified_time",
	OwnerInstanceID:                "owner_instance_id",
	PartnerHouseholdGraphDatasetID: "partner_household_graph_dataset_id",
	Region:                         "region",
	Status:                         "status",
	StudyCells:                     "study_cells",
	StudyEndTime:                   "study_end_time",
	StudyStartTime:                 "study_start_time",
}

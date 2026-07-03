package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"time"
)

type ThirdPartyPartnerLiftRequest struct {
	AdEntities                     *[]string                                 `json:"ad_entities,omitempty"`
	Country                        *string                                   `json:"country,omitempty"`
	CreatedTime                    *time.Time                                `json:"created_time,omitempty"`
	Description                    *string                                   `json:"description,omitempty"`
	HoldoutSize                    *float64                                  `json:"holdout_size,omitempty"`
	ID                             *core.ID                                  `json:"id,omitempty"`
	LegacyAdsDataPartnerID         *core.ID                                  `json:"legacy_ads_data_partner_id,omitempty"`
	LegacyAdsDataPartnerName       *string                                   `json:"legacy_ads_data_partner_name,omitempty"`
	ModifiedTime                   *time.Time                                `json:"modified_time,omitempty"`
	OwnerInstanceID                *core.ID                                  `json:"owner_instance_id,omitempty"`
	PartnerHouseholdGraphDatasetID *core.ID                                  `json:"partner_household_graph_dataset_id,omitempty"`
	Region                         *string                                   `json:"region,omitempty"`
	Status                         *enums.ThirdpartypartnerliftrequestStatus `json:"status,omitempty"`
	StudyCells                     *[]string                                 `json:"study_cells,omitempty"`
	StudyEndTime                   *time.Time                                `json:"study_end_time,omitempty"`
	StudyStartTime                 *time.Time                                `json:"study_start_time,omitempty"`
}

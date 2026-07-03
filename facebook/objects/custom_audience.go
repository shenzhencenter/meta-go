package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CustomAudience struct {
	AccountID                             *core.ID                      `json:"account_id,omitempty"`
	ApproximateCountLowerBound            *int                          `json:"approximate_count_lower_bound,omitempty"`
	ApproximateCountUpperBound            *int                          `json:"approximate_count_upper_bound,omitempty"`
	AudienceLabels                        *[]string                     `json:"audience_labels,omitempty"`
	CustomerFileSource                    *string                       `json:"customer_file_source,omitempty"`
	DataSource                            *CustomAudienceDataSource     `json:"data_source,omitempty"`
	DataSourceTypes                       *string                       `json:"data_source_types,omitempty"`
	DatafileCustomAudienceUploadingStatus *string                       `json:"datafile_custom_audience_uploading_status,omitempty"`
	DeleteTime                            *int                          `json:"delete_time,omitempty"`
	DeliveryStatus                        *CustomAudienceStatus         `json:"delivery_status,omitempty"`
	Description                           *string                       `json:"description,omitempty"`
	ExcludedCustomAudiences               *[]CustomAudience             `json:"excluded_custom_audiences,omitempty"`
	ExternalEventSource                   *AdsPixel                     `json:"external_event_source,omitempty"`
	FieldsViolatingIntegrityPolicy        *[]string                     `json:"fields_violating_integrity_policy,omitempty"`
	HouseholdAudience                     *int                          `json:"household_audience,omitempty"`
	ID                                    *core.ID                      `json:"id,omitempty"`
	IncludedCustomAudiences               *[]CustomAudience             `json:"included_custom_audiences,omitempty"`
	IsEligibleForSacCampaigns             *bool                         `json:"is_eligible_for_sac_campaigns,omitempty"`
	IsHousehold                           *bool                         `json:"is_household,omitempty"`
	IsSnapshot                            *bool                         `json:"is_snapshot,omitempty"`
	IsValueBased                          *bool                         `json:"is_value_based,omitempty"`
	LookalikeAudienceIds                  *[]core.ID                    `json:"lookalike_audience_ids,omitempty"`
	LookalikeSpec                         *LookalikeSpec                `json:"lookalike_spec,omitempty"`
	MessengerMarketingMessagesPage        *Page                         `json:"messenger_marketing_messages_page,omitempty"`
	Name                                  *string                       `json:"name,omitempty"`
	OperationStatus                       *CustomAudienceStatus         `json:"operation_status,omitempty"`
	OptOutLink                            *string                       `json:"opt_out_link,omitempty"`
	OwnerBusiness                         *Business                     `json:"owner_business,omitempty"`
	PageDeletionMarkedDeleteTime          *int                          `json:"page_deletion_marked_delete_time,omitempty"`
	PermissionForActions                  *AudiencePermissionForActions `json:"permission_for_actions,omitempty"`
	PixelID                               *core.ID                      `json:"pixel_id,omitempty"`
	RegulatedAudienceSpec                 *LookalikeSpec                `json:"regulated_audience_spec,omitempty"`
	RetentionDays                         *int                          `json:"retention_days,omitempty"`
	RevSharePolicyID                      *core.ID                      `json:"rev_share_policy_id,omitempty"`
	Rule                                  *string                       `json:"rule,omitempty"`
	RuleAggregation                       *string                       `json:"rule_aggregation,omitempty"`
	RuleV2                                *string                       `json:"rule_v2,omitempty"`
	SeedAudience                          *int                          `json:"seed_audience,omitempty"`
	SharingStatus                         *CustomAudienceSharingStatus  `json:"sharing_status,omitempty"`
	Subtype                               *string                       `json:"subtype,omitempty"`
	TimeContentUpdated                    *uint64                       `json:"time_content_updated,omitempty"`
	TimeCreated                           *uint64                       `json:"time_created,omitempty"`
	TimeUpdated                           *uint64                       `json:"time_updated,omitempty"`
}

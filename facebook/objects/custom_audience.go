package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var CustomAudienceFields = struct {
	AccountID                             string
	ApproximateCountLowerBound            string
	ApproximateCountUpperBound            string
	AudienceLabels                        string
	CustomerFileSource                    string
	DataSource                            string
	DataSourceTypes                       string
	DatafileCustomAudienceUploadingStatus string
	DeleteTime                            string
	DeliveryStatus                        string
	Description                           string
	ExcludedCustomAudiences               string
	ExternalEventSource                   string
	FieldsViolatingIntegrityPolicy        string
	HouseholdAudience                     string
	ID                                    string
	IncludedCustomAudiences               string
	IsEligibleForSacCampaigns             string
	IsHousehold                           string
	IsSnapshot                            string
	IsValueBased                          string
	LookalikeAudienceIds                  string
	LookalikeSpec                         string
	MessengerMarketingMessagesPage        string
	Name                                  string
	OperationStatus                       string
	OptOutLink                            string
	OwnerBusiness                         string
	PageDeletionMarkedDeleteTime          string
	PermissionForActions                  string
	PixelID                               string
	RegulatedAudienceSpec                 string
	RetentionDays                         string
	RevSharePolicyID                      string
	Rule                                  string
	RuleAggregation                       string
	RuleV2                                string
	SeedAudience                          string
	SharingStatus                         string
	Subtype                               string
	TimeContentUpdated                    string
	TimeCreated                           string
	TimeUpdated                           string
}{
	AccountID:                             "account_id",
	ApproximateCountLowerBound:            "approximate_count_lower_bound",
	ApproximateCountUpperBound:            "approximate_count_upper_bound",
	AudienceLabels:                        "audience_labels",
	CustomerFileSource:                    "customer_file_source",
	DataSource:                            "data_source",
	DataSourceTypes:                       "data_source_types",
	DatafileCustomAudienceUploadingStatus: "datafile_custom_audience_uploading_status",
	DeleteTime:                            "delete_time",
	DeliveryStatus:                        "delivery_status",
	Description:                           "description",
	ExcludedCustomAudiences:               "excluded_custom_audiences",
	ExternalEventSource:                   "external_event_source",
	FieldsViolatingIntegrityPolicy:        "fields_violating_integrity_policy",
	HouseholdAudience:                     "household_audience",
	ID:                                    "id",
	IncludedCustomAudiences:               "included_custom_audiences",
	IsEligibleForSacCampaigns:             "is_eligible_for_sac_campaigns",
	IsHousehold:                           "is_household",
	IsSnapshot:                            "is_snapshot",
	IsValueBased:                          "is_value_based",
	LookalikeAudienceIds:                  "lookalike_audience_ids",
	LookalikeSpec:                         "lookalike_spec",
	MessengerMarketingMessagesPage:        "messenger_marketing_messages_page",
	Name:                                  "name",
	OperationStatus:                       "operation_status",
	OptOutLink:                            "opt_out_link",
	OwnerBusiness:                         "owner_business",
	PageDeletionMarkedDeleteTime:          "page_deletion_marked_delete_time",
	PermissionForActions:                  "permission_for_actions",
	PixelID:                               "pixel_id",
	RegulatedAudienceSpec:                 "regulated_audience_spec",
	RetentionDays:                         "retention_days",
	RevSharePolicyID:                      "rev_share_policy_id",
	Rule:                                  "rule",
	RuleAggregation:                       "rule_aggregation",
	RuleV2:                                "rule_v2",
	SeedAudience:                          "seed_audience",
	SharingStatus:                         "sharing_status",
	Subtype:                               "subtype",
	TimeContentUpdated:                    "time_content_updated",
	TimeCreated:                           "time_created",
	TimeUpdated:                           "time_updated",
}

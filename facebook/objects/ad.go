package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type Ad struct {
	AccountID                         *core.ID                           `json:"account_id,omitempty"`
	AdActiveTime                      *string                            `json:"ad_active_time,omitempty"`
	AdReviewFeedback                  *AdgroupReviewFeedback             `json:"ad_review_feedback,omitempty"`
	AdScheduleEndTime                 *core.Time                         `json:"ad_schedule_end_time,omitempty"`
	AdScheduleStartTime               *core.Time                         `json:"ad_schedule_start_time,omitempty"`
	Adlabels                          *[]AdLabel                         `json:"adlabels,omitempty"`
	Adset                             *AdSet                             `json:"adset,omitempty"`
	AdsetID                           *core.ID                           `json:"adset_id,omitempty"`
	BidAmount                         *int                               `json:"bid_amount,omitempty"`
	BidInfo                           *map[string]uint64                 `json:"bid_info,omitempty"`
	BidType                           *enums.AdBidType                   `json:"bid_type,omitempty"`
	Campaign                          *Campaign                          `json:"campaign,omitempty"`
	CampaignID                        *core.ID                           `json:"campaign_id,omitempty"`
	ConfiguredStatus                  *enums.AdConfiguredStatus          `json:"configured_status,omitempty"`
	ConversionDomain                  *string                            `json:"conversion_domain,omitempty"`
	ConversionSpecs                   *[]ConversionActionQuery           `json:"conversion_specs,omitempty"`
	CreatedTime                       *core.Time                         `json:"created_time,omitempty"`
	Creative                          *AdCreative                        `json:"creative,omitempty"`
	CreativeAssetGroupsSpec           *AdCreativeAssetGroupsSpec         `json:"creative_asset_groups_spec,omitempty"`
	CreativeAutomationSpec            *AdCreativeAutomationSpec          `json:"creative_automation_spec,omitempty"`
	DemolinkHash                      *string                            `json:"demolink_hash,omitempty"`
	DisplaySequence                   *int                               `json:"display_sequence,omitempty"`
	EffectiveStatus                   *enums.AdEffectiveStatus           `json:"effective_status,omitempty"`
	EngagementAudience                *bool                              `json:"engagement_audience,omitempty"`
	FailedDeliveryChecks              *[]DeliveryCheck                   `json:"failed_delivery_checks,omitempty"`
	ID                                *core.ID                           `json:"id,omitempty"`
	IssuesInfo                        *[]AdgroupIssuesInfo               `json:"issues_info,omitempty"`
	LastUpdatedByAppID                *core.ID                           `json:"last_updated_by_app_id,omitempty"`
	Name                              *string                            `json:"name,omitempty"`
	Placement                         *Placement                         `json:"placement,omitempty"`
	PreviewShareableLink              *string                            `json:"preview_shareable_link,omitempty"`
	Priority                          *uint64                            `json:"priority,omitempty"`
	Recommendations                   *[]AdRecommendation                `json:"recommendations,omitempty"`
	SourceAd                          *Ad                                `json:"source_ad,omitempty"`
	SourceAdID                        *core.ID                           `json:"source_ad_id,omitempty"`
	SpecialAdCategories               *[]string                          `json:"special_ad_categories,omitempty"`
	Status                            *enums.AdStatus                    `json:"status,omitempty"`
	Targeting                         *Targeting                         `json:"targeting,omitempty"`
	TrackingAndConversionWithDefaults *TrackingAndConversionWithDefaults `json:"tracking_and_conversion_with_defaults,omitempty"`
	TrackingSpecs                     *[]ConversionActionQuery           `json:"tracking_specs,omitempty"`
	UpdatedTime                       *core.Time                         `json:"updated_time,omitempty"`
}

var AdFields = struct {
	AccountID                         string
	AdActiveTime                      string
	AdReviewFeedback                  string
	AdScheduleEndTime                 string
	AdScheduleStartTime               string
	Adlabels                          string
	Adset                             string
	AdsetID                           string
	BidAmount                         string
	BidInfo                           string
	BidType                           string
	Campaign                          string
	CampaignID                        string
	ConfiguredStatus                  string
	ConversionDomain                  string
	ConversionSpecs                   string
	CreatedTime                       string
	Creative                          string
	CreativeAssetGroupsSpec           string
	CreativeAutomationSpec            string
	DemolinkHash                      string
	DisplaySequence                   string
	EffectiveStatus                   string
	EngagementAudience                string
	FailedDeliveryChecks              string
	ID                                string
	IssuesInfo                        string
	LastUpdatedByAppID                string
	Name                              string
	Placement                         string
	PreviewShareableLink              string
	Priority                          string
	Recommendations                   string
	SourceAd                          string
	SourceAdID                        string
	SpecialAdCategories               string
	Status                            string
	Targeting                         string
	TrackingAndConversionWithDefaults string
	TrackingSpecs                     string
	UpdatedTime                       string
}{
	AccountID:                         "account_id",
	AdActiveTime:                      "ad_active_time",
	AdReviewFeedback:                  "ad_review_feedback",
	AdScheduleEndTime:                 "ad_schedule_end_time",
	AdScheduleStartTime:               "ad_schedule_start_time",
	Adlabels:                          "adlabels",
	Adset:                             "adset",
	AdsetID:                           "adset_id",
	BidAmount:                         "bid_amount",
	BidInfo:                           "bid_info",
	BidType:                           "bid_type",
	Campaign:                          "campaign",
	CampaignID:                        "campaign_id",
	ConfiguredStatus:                  "configured_status",
	ConversionDomain:                  "conversion_domain",
	ConversionSpecs:                   "conversion_specs",
	CreatedTime:                       "created_time",
	Creative:                          "creative",
	CreativeAssetGroupsSpec:           "creative_asset_groups_spec",
	CreativeAutomationSpec:            "creative_automation_spec",
	DemolinkHash:                      "demolink_hash",
	DisplaySequence:                   "display_sequence",
	EffectiveStatus:                   "effective_status",
	EngagementAudience:                "engagement_audience",
	FailedDeliveryChecks:              "failed_delivery_checks",
	ID:                                "id",
	IssuesInfo:                        "issues_info",
	LastUpdatedByAppID:                "last_updated_by_app_id",
	Name:                              "name",
	Placement:                         "placement",
	PreviewShareableLink:              "preview_shareable_link",
	Priority:                          "priority",
	Recommendations:                   "recommendations",
	SourceAd:                          "source_ad",
	SourceAdID:                        "source_ad_id",
	SpecialAdCategories:               "special_ad_categories",
	Status:                            "status",
	Targeting:                         "targeting",
	TrackingAndConversionWithDefaults: "tracking_and_conversion_with_defaults",
	TrackingSpecs:                     "tracking_specs",
	UpdatedTime:                       "updated_time",
}

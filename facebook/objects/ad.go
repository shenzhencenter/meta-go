package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"time"
)

type Ad struct {
	AccountID                         *core.ID                           `json:"account_id,omitempty"`
	AdActiveTime                      *string                            `json:"ad_active_time,omitempty"`
	AdReviewFeedback                  *AdgroupReviewFeedback             `json:"ad_review_feedback,omitempty"`
	AdScheduleEndTime                 *time.Time                         `json:"ad_schedule_end_time,omitempty"`
	AdScheduleStartTime               *time.Time                         `json:"ad_schedule_start_time,omitempty"`
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
	CreatedTime                       *time.Time                         `json:"created_time,omitempty"`
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
	UpdatedTime                       *time.Time                         `json:"updated_time,omitempty"`
}

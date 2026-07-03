package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Lead struct {
	AdID                      *core.ID                          `json:"ad_id,omitempty"`
	AdName                    *string                           `json:"ad_name,omitempty"`
	AdsetID                   *core.ID                          `json:"adset_id,omitempty"`
	AdsetName                 *string                           `json:"adset_name,omitempty"`
	CampaignID                *core.ID                          `json:"campaign_id,omitempty"`
	CampaignName              *string                           `json:"campaign_name,omitempty"`
	CreatedTime               *core.Time                        `json:"created_time,omitempty"`
	CustomDisclaimerResponses *[]UserLeadGenDisclaimerResponse  `json:"custom_disclaimer_responses,omitempty"`
	FieldData                 *[]UserLeadGenFieldData           `json:"field_data,omitempty"`
	FormID                    *core.ID                          `json:"form_id,omitempty"`
	HomeListing               *HomeListing                      `json:"home_listing,omitempty"`
	ID                        *core.ID                          `json:"id,omitempty"`
	IsOrganic                 *bool                             `json:"is_organic,omitempty"`
	PartnerName               *string                           `json:"partner_name,omitempty"`
	Platform                  *string                           `json:"platform,omitempty"`
	Post                      *Link                             `json:"post,omitempty"`
	PostSubmissionCheckResult *LeadGenPostSubmissionCheckResult `json:"post_submission_check_result,omitempty"`
	RetailerItemID            *core.ID                          `json:"retailer_item_id,omitempty"`
	Vehicle                   *Vehicle                          `json:"vehicle,omitempty"`
}

var LeadFields = struct {
	AdID                      string
	AdName                    string
	AdsetID                   string
	AdsetName                 string
	CampaignID                string
	CampaignName              string
	CreatedTime               string
	CustomDisclaimerResponses string
	FieldData                 string
	FormID                    string
	HomeListing               string
	ID                        string
	IsOrganic                 string
	PartnerName               string
	Platform                  string
	Post                      string
	PostSubmissionCheckResult string
	RetailerItemID            string
	Vehicle                   string
}{
	AdID:                      "ad_id",
	AdName:                    "ad_name",
	AdsetID:                   "adset_id",
	AdsetName:                 "adset_name",
	CampaignID:                "campaign_id",
	CampaignName:              "campaign_name",
	CreatedTime:               "created_time",
	CustomDisclaimerResponses: "custom_disclaimer_responses",
	FieldData:                 "field_data",
	FormID:                    "form_id",
	HomeListing:               "home_listing",
	ID:                        "id",
	IsOrganic:                 "is_organic",
	PartnerName:               "partner_name",
	Platform:                  "platform",
	Post:                      "post",
	PostSubmissionCheckResult: "post_submission_check_result",
	RetailerItemID:            "retailer_item_id",
	Vehicle:                   "vehicle",
}

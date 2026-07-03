package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type LeadgenForm struct {
	AllowOrganicLead                 *bool                `json:"allow_organic_lead,omitempty"`
	BlockDisplayForNonTargetedViewer *bool                `json:"block_display_for_non_targeted_viewer,omitempty"`
	ContextCard                      *LeadGenContextCard  `json:"context_card,omitempty"`
	CreatedTime                      *core.Time           `json:"created_time,omitempty"`
	Creator                          *User                `json:"creator,omitempty"`
	ExpiredLeadsCount                *uint64              `json:"expired_leads_count,omitempty"`
	FollowUpActionText               *string              `json:"follow_up_action_text,omitempty"`
	FollowUpActionURL                *string              `json:"follow_up_action_url,omitempty"`
	ID                               *core.ID             `json:"id,omitempty"`
	IsOptimizedForQuality            *bool                `json:"is_optimized_for_quality,omitempty"`
	LeadsCount                       *uint64              `json:"leads_count,omitempty"`
	LegalContent                     *LeadGenLegalContent `json:"legal_content,omitempty"`
	Locale                           *string              `json:"locale,omitempty"`
	Name                             *string              `json:"name,omitempty"`
	OrganicLeadsCount                *uint64              `json:"organic_leads_count,omitempty"`
	Page                             *Page                `json:"page,omitempty"`
	PageID                           *core.ID             `json:"page_id,omitempty"`
	PrivacyPolicyURL                 *string              `json:"privacy_policy_url,omitempty"`
	QuestionPageCustomHeadline       *string              `json:"question_page_custom_headline,omitempty"`
	Questions                        *[]LeadGenQuestion   `json:"questions,omitempty"`
	Status                           *string              `json:"status,omitempty"`
	ThankYouPage                     *LeadGenThankYouPage `json:"thank_you_page,omitempty"`
	TrackingParameters               *[]map[string]string `json:"tracking_parameters,omitempty"`
}

var LeadgenFormFields = struct {
	AllowOrganicLead                 string
	BlockDisplayForNonTargetedViewer string
	ContextCard                      string
	CreatedTime                      string
	Creator                          string
	ExpiredLeadsCount                string
	FollowUpActionText               string
	FollowUpActionURL                string
	ID                               string
	IsOptimizedForQuality            string
	LeadsCount                       string
	LegalContent                     string
	Locale                           string
	Name                             string
	OrganicLeadsCount                string
	Page                             string
	PageID                           string
	PrivacyPolicyURL                 string
	QuestionPageCustomHeadline       string
	Questions                        string
	Status                           string
	ThankYouPage                     string
	TrackingParameters               string
}{
	AllowOrganicLead:                 "allow_organic_lead",
	BlockDisplayForNonTargetedViewer: "block_display_for_non_targeted_viewer",
	ContextCard:                      "context_card",
	CreatedTime:                      "created_time",
	Creator:                          "creator",
	ExpiredLeadsCount:                "expired_leads_count",
	FollowUpActionText:               "follow_up_action_text",
	FollowUpActionURL:                "follow_up_action_url",
	ID:                               "id",
	IsOptimizedForQuality:            "is_optimized_for_quality",
	LeadsCount:                       "leads_count",
	LegalContent:                     "legal_content",
	Locale:                           "locale",
	Name:                             "name",
	OrganicLeadsCount:                "organic_leads_count",
	Page:                             "page",
	PageID:                           "page_id",
	PrivacyPolicyURL:                 "privacy_policy_url",
	QuestionPageCustomHeadline:       "question_page_custom_headline",
	Questions:                        "questions",
	Status:                           "status",
	ThankYouPage:                     "thank_you_page",
	TrackingParameters:               "tracking_parameters",
}

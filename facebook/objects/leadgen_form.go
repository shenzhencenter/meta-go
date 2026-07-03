package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type LeadgenForm struct {
	AllowOrganicLead                 *bool                `json:"allow_organic_lead,omitempty"`
	BlockDisplayForNonTargetedViewer *bool                `json:"block_display_for_non_targeted_viewer,omitempty"`
	ContextCard                      *LeadGenContextCard  `json:"context_card,omitempty"`
	CreatedTime                      *time.Time           `json:"created_time,omitempty"`
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

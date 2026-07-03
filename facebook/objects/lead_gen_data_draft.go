package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type LeadGenDataDraft struct {
	BlockDisplayForNonTargetedViewer *bool                   `json:"block_display_for_non_targeted_viewer,omitempty"`
	CreatedTime                      *core.Time              `json:"created_time,omitempty"`
	DisqualifiedEndComponent         *map[string]interface{} `json:"disqualified_end_component,omitempty"`
	FollowUpActionURL                *string                 `json:"follow_up_action_url,omitempty"`
	ID                               *core.ID                `json:"id,omitempty"`
	IsOptimizedForQuality            *bool                   `json:"is_optimized_for_quality,omitempty"`
	LegalContent                     *map[string]interface{} `json:"legal_content,omitempty"`
	Locale                           *string                 `json:"locale,omitempty"`
	Name                             *string                 `json:"name,omitempty"`
	Page                             *Page                   `json:"page,omitempty"`
	QuestionPageCustomHeadline       *string                 `json:"question_page_custom_headline,omitempty"`
	Questions                        *[]LeadGenDraftQuestion `json:"questions,omitempty"`
	ShouldEnforceWorkEmail           *bool                   `json:"should_enforce_work_email,omitempty"`
	Status                           *string                 `json:"status,omitempty"`
	ThankYouPage                     *map[string]interface{} `json:"thank_you_page,omitempty"`
	TrackingParameters               *[]map[string]string    `json:"tracking_parameters,omitempty"`
}

var LeadGenDataDraftFields = struct {
	BlockDisplayForNonTargetedViewer string
	CreatedTime                      string
	DisqualifiedEndComponent         string
	FollowUpActionURL                string
	ID                               string
	IsOptimizedForQuality            string
	LegalContent                     string
	Locale                           string
	Name                             string
	Page                             string
	QuestionPageCustomHeadline       string
	Questions                        string
	ShouldEnforceWorkEmail           string
	Status                           string
	ThankYouPage                     string
	TrackingParameters               string
}{
	BlockDisplayForNonTargetedViewer: "block_display_for_non_targeted_viewer",
	CreatedTime:                      "created_time",
	DisqualifiedEndComponent:         "disqualified_end_component",
	FollowUpActionURL:                "follow_up_action_url",
	ID:                               "id",
	IsOptimizedForQuality:            "is_optimized_for_quality",
	LegalContent:                     "legal_content",
	Locale:                           "locale",
	Name:                             "name",
	Page:                             "page",
	QuestionPageCustomHeadline:       "question_page_custom_headline",
	Questions:                        "questions",
	ShouldEnforceWorkEmail:           "should_enforce_work_email",
	Status:                           "status",
	ThankYouPage:                     "thank_you_page",
	TrackingParameters:               "tracking_parameters",
}

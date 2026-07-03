package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type LeadGenDataDraft struct {
	BlockDisplayForNonTargetedViewer *bool                   `json:"block_display_for_non_targeted_viewer,omitempty"`
	CreatedTime                      *time.Time              `json:"created_time,omitempty"`
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

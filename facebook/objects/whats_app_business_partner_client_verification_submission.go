package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type WhatsAppBusinessPartnerClientVerificationSubmission struct {
	ClientBusinessID   *core.ID                                                                     `json:"client_business_id,omitempty"`
	ID                 *core.ID                                                                     `json:"id,omitempty"`
	RejectionReasons   *[]enums.WhatsappbusinesspartnerclientverificationsubmissionRejectionReasons `json:"rejection_reasons,omitempty"`
	SubmittedInfo      *map[string]interface{}                                                      `json:"submitted_info,omitempty"`
	SubmittedTime      *core.Time                                                                   `json:"submitted_time,omitempty"`
	UpdateTime         *core.Time                                                                   `json:"update_time,omitempty"`
	VerificationStatus *enums.WhatsappbusinesspartnerclientverificationsubmissionVerificationStatus `json:"verification_status,omitempty"`
}

var WhatsAppBusinessPartnerClientVerificationSubmissionFields = struct {
	ClientBusinessID   string
	ID                 string
	RejectionReasons   string
	SubmittedInfo      string
	SubmittedTime      string
	UpdateTime         string
	VerificationStatus string
}{
	ClientBusinessID:   "client_business_id",
	ID:                 "id",
	RejectionReasons:   "rejection_reasons",
	SubmittedInfo:      "submitted_info",
	SubmittedTime:      "submitted_time",
	UpdateTime:         "update_time",
	VerificationStatus: "verification_status",
}

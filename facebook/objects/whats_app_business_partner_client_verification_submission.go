package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"time"
)

type WhatsAppBusinessPartnerClientVerificationSubmission struct {
	ClientBusinessID   *core.ID                                                                     `json:"client_business_id,omitempty"`
	ID                 *core.ID                                                                     `json:"id,omitempty"`
	RejectionReasons   *[]enums.WhatsappbusinesspartnerclientverificationsubmissionRejectionReasons `json:"rejection_reasons,omitempty"`
	SubmittedInfo      *map[string]interface{}                                                      `json:"submitted_info,omitempty"`
	SubmittedTime      *time.Time                                                                   `json:"submitted_time,omitempty"`
	UpdateTime         *time.Time                                                                   `json:"update_time,omitempty"`
	VerificationStatus *enums.WhatsappbusinesspartnerclientverificationsubmissionVerificationStatus `json:"verification_status,omitempty"`
}

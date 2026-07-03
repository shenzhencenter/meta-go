package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"time"
)

type WhatsAppBusinessPreVerifiedPhoneNumber struct {
	CodeVerificationStatus *enums.WhatsappbusinesspreverifiedphonenumberCodeVerificationStatus `json:"code_verification_status,omitempty"`
	CodeVerificationTime   *time.Time                                                          `json:"code_verification_time,omitempty"`
	ID                     *core.ID                                                            `json:"id,omitempty"`
	OwnerBusiness          *Business                                                           `json:"owner_business,omitempty"`
	PhoneNumber            *string                                                             `json:"phone_number,omitempty"`
	VerificationExpiryTime *time.Time                                                          `json:"verification_expiry_time,omitempty"`
}

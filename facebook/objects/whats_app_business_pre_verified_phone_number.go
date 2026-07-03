package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type WhatsAppBusinessPreVerifiedPhoneNumber struct {
	CodeVerificationStatus *enums.WhatsappbusinesspreverifiedphonenumberCodeVerificationStatus `json:"code_verification_status,omitempty"`
	CodeVerificationTime   *core.Time                                                          `json:"code_verification_time,omitempty"`
	ID                     *core.ID                                                            `json:"id,omitempty"`
	OwnerBusiness          *Business                                                           `json:"owner_business,omitempty"`
	PhoneNumber            *string                                                             `json:"phone_number,omitempty"`
	VerificationExpiryTime *core.Time                                                          `json:"verification_expiry_time,omitempty"`
}

var WhatsAppBusinessPreVerifiedPhoneNumberFields = struct {
	CodeVerificationStatus string
	CodeVerificationTime   string
	ID                     string
	OwnerBusiness          string
	PhoneNumber            string
	VerificationExpiryTime string
}{
	CodeVerificationStatus: "code_verification_status",
	CodeVerificationTime:   "code_verification_time",
	ID:                     "id",
	OwnerBusiness:          "owner_business",
	PhoneNumber:            "phone_number",
	VerificationExpiryTime: "verification_expiry_time",
}

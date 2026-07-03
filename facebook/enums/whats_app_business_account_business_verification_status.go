package enums

type WhatsappbusinessaccountBusinessVerificationStatus string

const (
	WhatsappbusinessaccountBusinessVerificationStatusExpired             WhatsappbusinessaccountBusinessVerificationStatus = "expired"
	WhatsappbusinessaccountBusinessVerificationStatusFailed              WhatsappbusinessaccountBusinessVerificationStatus = "failed"
	WhatsappbusinessaccountBusinessVerificationStatusIneligible          WhatsappbusinessaccountBusinessVerificationStatus = "ineligible"
	WhatsappbusinessaccountBusinessVerificationStatusNotVerified         WhatsappbusinessaccountBusinessVerificationStatus = "not_verified"
	WhatsappbusinessaccountBusinessVerificationStatusPending             WhatsappbusinessaccountBusinessVerificationStatus = "pending"
	WhatsappbusinessaccountBusinessVerificationStatusPendingNeedMoreInfo WhatsappbusinessaccountBusinessVerificationStatus = "pending_need_more_info"
	WhatsappbusinessaccountBusinessVerificationStatusPendingSubmission   WhatsappbusinessaccountBusinessVerificationStatus = "pending_submission"
	WhatsappbusinessaccountBusinessVerificationStatusRejected            WhatsappbusinessaccountBusinessVerificationStatus = "rejected"
	WhatsappbusinessaccountBusinessVerificationStatusRevoked             WhatsappbusinessaccountBusinessVerificationStatus = "revoked"
	WhatsappbusinessaccountBusinessVerificationStatusVerified            WhatsappbusinessaccountBusinessVerificationStatus = "verified"
)

func (value WhatsappbusinessaccountBusinessVerificationStatus) String() string {
	return string(value)
}

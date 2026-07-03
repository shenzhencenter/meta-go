package enums

type BusinessVerificationStatus string

const (
	BusinessVerificationStatusExpired             BusinessVerificationStatus = "expired"
	BusinessVerificationStatusFailed              BusinessVerificationStatus = "failed"
	BusinessVerificationStatusIneligible          BusinessVerificationStatus = "ineligible"
	BusinessVerificationStatusNotVerified         BusinessVerificationStatus = "not_verified"
	BusinessVerificationStatusPending             BusinessVerificationStatus = "pending"
	BusinessVerificationStatusPendingNeedMoreInfo BusinessVerificationStatus = "pending_need_more_info"
	BusinessVerificationStatusPendingSubmission   BusinessVerificationStatus = "pending_submission"
	BusinessVerificationStatusRejected            BusinessVerificationStatus = "rejected"
	BusinessVerificationStatusRevoked             BusinessVerificationStatus = "revoked"
	BusinessVerificationStatusVerified            BusinessVerificationStatus = "verified"
)

func (value BusinessVerificationStatus) String() string {
	return string(value)
}

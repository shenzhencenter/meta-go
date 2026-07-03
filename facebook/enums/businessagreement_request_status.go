package enums

type BusinessagreementRequestStatus string

const (
	BusinessagreementRequestStatusApprove                  BusinessagreementRequestStatus = "APPROVE"
	BusinessagreementRequestStatusCanceled                 BusinessagreementRequestStatus = "CANCELED"
	BusinessagreementRequestStatusDecline                  BusinessagreementRequestStatus = "DECLINE"
	BusinessagreementRequestStatusExpired                  BusinessagreementRequestStatus = "EXPIRED"
	BusinessagreementRequestStatusInProgress               BusinessagreementRequestStatus = "IN_PROGRESS"
	BusinessagreementRequestStatusMmaDirectAssetsApproved  BusinessagreementRequestStatus = "MMA_DIRECT_ASSETS_APPROVED"
	BusinessagreementRequestStatusMmaDirectAssetsDeclined  BusinessagreementRequestStatus = "MMA_DIRECT_ASSETS_DECLINED"
	BusinessagreementRequestStatusMmaDirectAssetsExpired   BusinessagreementRequestStatus = "MMA_DIRECT_ASSETS_EXPIRED"
	BusinessagreementRequestStatusMmaDirectAssetsPending   BusinessagreementRequestStatus = "MMA_DIRECT_ASSETS_PENDING"
	BusinessagreementRequestStatusPending                  BusinessagreementRequestStatus = "PENDING"
	BusinessagreementRequestStatusPendingEmailVerification BusinessagreementRequestStatus = "PENDING_EMAIL_VERIFICATION"
	BusinessagreementRequestStatusPendingIntegrityReview   BusinessagreementRequestStatus = "PENDING_INTEGRITY_REVIEW"
)

func (value BusinessagreementRequestStatus) String() string {
	return string(value)
}

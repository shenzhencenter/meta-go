package enums

type AdaccountonbehalfRequestsStatusEnumParam string

const (
	AdaccountonbehalfRequestsStatusEnumParamApprove                  AdaccountonbehalfRequestsStatusEnumParam = "APPROVE"
	AdaccountonbehalfRequestsStatusEnumParamCanceled                 AdaccountonbehalfRequestsStatusEnumParam = "CANCELED"
	AdaccountonbehalfRequestsStatusEnumParamDecline                  AdaccountonbehalfRequestsStatusEnumParam = "DECLINE"
	AdaccountonbehalfRequestsStatusEnumParamExpired                  AdaccountonbehalfRequestsStatusEnumParam = "EXPIRED"
	AdaccountonbehalfRequestsStatusEnumParamInProgress               AdaccountonbehalfRequestsStatusEnumParam = "IN_PROGRESS"
	AdaccountonbehalfRequestsStatusEnumParamMmaDirectAssetsApproved  AdaccountonbehalfRequestsStatusEnumParam = "MMA_DIRECT_ASSETS_APPROVED"
	AdaccountonbehalfRequestsStatusEnumParamMmaDirectAssetsDeclined  AdaccountonbehalfRequestsStatusEnumParam = "MMA_DIRECT_ASSETS_DECLINED"
	AdaccountonbehalfRequestsStatusEnumParamMmaDirectAssetsExpired   AdaccountonbehalfRequestsStatusEnumParam = "MMA_DIRECT_ASSETS_EXPIRED"
	AdaccountonbehalfRequestsStatusEnumParamMmaDirectAssetsPending   AdaccountonbehalfRequestsStatusEnumParam = "MMA_DIRECT_ASSETS_PENDING"
	AdaccountonbehalfRequestsStatusEnumParamPending                  AdaccountonbehalfRequestsStatusEnumParam = "PENDING"
	AdaccountonbehalfRequestsStatusEnumParamPendingEmailVerification AdaccountonbehalfRequestsStatusEnumParam = "PENDING_EMAIL_VERIFICATION"
	AdaccountonbehalfRequestsStatusEnumParamPendingIntegrityReview   AdaccountonbehalfRequestsStatusEnumParam = "PENDING_INTEGRITY_REVIEW"
)

func (value AdaccountonbehalfRequestsStatusEnumParam) String() string {
	return string(value)
}

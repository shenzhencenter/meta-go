package enums

type BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParam string

const (
	BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParamApprove                  BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParam = "APPROVE"
	BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParamCanceled                 BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParam = "CANCELED"
	BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParamDecline                  BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParam = "DECLINE"
	BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParamExpired                  BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParam = "EXPIRED"
	BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParamInProgress               BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParam = "IN_PROGRESS"
	BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParamMmaDirectAssetsApproved  BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParam = "MMA_DIRECT_ASSETS_APPROVED"
	BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParamMmaDirectAssetsDeclined  BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParam = "MMA_DIRECT_ASSETS_DECLINED"
	BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParamMmaDirectAssetsExpired   BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParam = "MMA_DIRECT_ASSETS_EXPIRED"
	BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParamMmaDirectAssetsPending   BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParam = "MMA_DIRECT_ASSETS_PENDING"
	BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParamPending                  BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParam = "PENDING"
	BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParamPendingEmailVerification BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParam = "PENDING_EMAIL_VERIFICATION"
	BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParamPendingIntegrityReview   BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParam = "PENDING_INTEGRITY_REVIEW"
)

func (value BusinessinitiatedAudienceSharingRequestsRequestStatusEnumParam) String() string {
	return string(value)
}

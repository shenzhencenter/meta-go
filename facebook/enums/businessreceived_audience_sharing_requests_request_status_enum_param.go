package enums

type BusinessreceivedAudienceSharingRequestsRequestStatusEnumParam string

const (
	BusinessreceivedAudienceSharingRequestsRequestStatusEnumParamApprove                  BusinessreceivedAudienceSharingRequestsRequestStatusEnumParam = "APPROVE"
	BusinessreceivedAudienceSharingRequestsRequestStatusEnumParamCanceled                 BusinessreceivedAudienceSharingRequestsRequestStatusEnumParam = "CANCELED"
	BusinessreceivedAudienceSharingRequestsRequestStatusEnumParamDecline                  BusinessreceivedAudienceSharingRequestsRequestStatusEnumParam = "DECLINE"
	BusinessreceivedAudienceSharingRequestsRequestStatusEnumParamExpired                  BusinessreceivedAudienceSharingRequestsRequestStatusEnumParam = "EXPIRED"
	BusinessreceivedAudienceSharingRequestsRequestStatusEnumParamInProgress               BusinessreceivedAudienceSharingRequestsRequestStatusEnumParam = "IN_PROGRESS"
	BusinessreceivedAudienceSharingRequestsRequestStatusEnumParamMmaDirectAssetsApproved  BusinessreceivedAudienceSharingRequestsRequestStatusEnumParam = "MMA_DIRECT_ASSETS_APPROVED"
	BusinessreceivedAudienceSharingRequestsRequestStatusEnumParamMmaDirectAssetsDeclined  BusinessreceivedAudienceSharingRequestsRequestStatusEnumParam = "MMA_DIRECT_ASSETS_DECLINED"
	BusinessreceivedAudienceSharingRequestsRequestStatusEnumParamMmaDirectAssetsExpired   BusinessreceivedAudienceSharingRequestsRequestStatusEnumParam = "MMA_DIRECT_ASSETS_EXPIRED"
	BusinessreceivedAudienceSharingRequestsRequestStatusEnumParamMmaDirectAssetsPending   BusinessreceivedAudienceSharingRequestsRequestStatusEnumParam = "MMA_DIRECT_ASSETS_PENDING"
	BusinessreceivedAudienceSharingRequestsRequestStatusEnumParamPending                  BusinessreceivedAudienceSharingRequestsRequestStatusEnumParam = "PENDING"
	BusinessreceivedAudienceSharingRequestsRequestStatusEnumParamPendingEmailVerification BusinessreceivedAudienceSharingRequestsRequestStatusEnumParam = "PENDING_EMAIL_VERIFICATION"
	BusinessreceivedAudienceSharingRequestsRequestStatusEnumParamPendingIntegrityReview   BusinessreceivedAudienceSharingRequestsRequestStatusEnumParam = "PENDING_INTEGRITY_REVIEW"
)

func (value BusinessreceivedAudienceSharingRequestsRequestStatusEnumParam) String() string {
	return string(value)
}

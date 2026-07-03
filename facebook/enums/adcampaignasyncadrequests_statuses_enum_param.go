package enums

type AdcampaignasyncadrequestsStatusesEnumParam string

const (
	AdcampaignasyncadrequestsStatusesEnumParamCanceled                AdcampaignasyncadrequestsStatusesEnumParam = "CANCELED"
	AdcampaignasyncadrequestsStatusesEnumParamCanceledDependency      AdcampaignasyncadrequestsStatusesEnumParam = "CANCELED_DEPENDENCY"
	AdcampaignasyncadrequestsStatusesEnumParamError                   AdcampaignasyncadrequestsStatusesEnumParam = "ERROR"
	AdcampaignasyncadrequestsStatusesEnumParamErrorConflicts          AdcampaignasyncadrequestsStatusesEnumParam = "ERROR_CONFLICTS"
	AdcampaignasyncadrequestsStatusesEnumParamErrorDependency         AdcampaignasyncadrequestsStatusesEnumParam = "ERROR_DEPENDENCY"
	AdcampaignasyncadrequestsStatusesEnumParamInitial                 AdcampaignasyncadrequestsStatusesEnumParam = "INITIAL"
	AdcampaignasyncadrequestsStatusesEnumParamInProgress              AdcampaignasyncadrequestsStatusesEnumParam = "IN_PROGRESS"
	AdcampaignasyncadrequestsStatusesEnumParamPendingDependency       AdcampaignasyncadrequestsStatusesEnumParam = "PENDING_DEPENDENCY"
	AdcampaignasyncadrequestsStatusesEnumParamProcessByAdAsyncEngine  AdcampaignasyncadrequestsStatusesEnumParam = "PROCESS_BY_AD_ASYNC_ENGINE"
	AdcampaignasyncadrequestsStatusesEnumParamProcessByEventProcessor AdcampaignasyncadrequestsStatusesEnumParam = "PROCESS_BY_EVENT_PROCESSOR"
	AdcampaignasyncadrequestsStatusesEnumParamSuccess                 AdcampaignasyncadrequestsStatusesEnumParam = "SUCCESS"
	AdcampaignasyncadrequestsStatusesEnumParamUserCanceled            AdcampaignasyncadrequestsStatusesEnumParam = "USER_CANCELED"
	AdcampaignasyncadrequestsStatusesEnumParamUserCanceledDependency  AdcampaignasyncadrequestsStatusesEnumParam = "USER_CANCELED_DEPENDENCY"
)

func (value AdcampaignasyncadrequestsStatusesEnumParam) String() string {
	return string(value)
}

package enums

type AdasyncrequestsetrequestsStatusesEnumParam string

const (
	AdasyncrequestsetrequestsStatusesEnumParamCanceled                AdasyncrequestsetrequestsStatusesEnumParam = "CANCELED"
	AdasyncrequestsetrequestsStatusesEnumParamCanceledDependency      AdasyncrequestsetrequestsStatusesEnumParam = "CANCELED_DEPENDENCY"
	AdasyncrequestsetrequestsStatusesEnumParamError                   AdasyncrequestsetrequestsStatusesEnumParam = "ERROR"
	AdasyncrequestsetrequestsStatusesEnumParamErrorConflicts          AdasyncrequestsetrequestsStatusesEnumParam = "ERROR_CONFLICTS"
	AdasyncrequestsetrequestsStatusesEnumParamErrorDependency         AdasyncrequestsetrequestsStatusesEnumParam = "ERROR_DEPENDENCY"
	AdasyncrequestsetrequestsStatusesEnumParamInitial                 AdasyncrequestsetrequestsStatusesEnumParam = "INITIAL"
	AdasyncrequestsetrequestsStatusesEnumParamInProgress              AdasyncrequestsetrequestsStatusesEnumParam = "IN_PROGRESS"
	AdasyncrequestsetrequestsStatusesEnumParamPendingDependency       AdasyncrequestsetrequestsStatusesEnumParam = "PENDING_DEPENDENCY"
	AdasyncrequestsetrequestsStatusesEnumParamProcessByAdAsyncEngine  AdasyncrequestsetrequestsStatusesEnumParam = "PROCESS_BY_AD_ASYNC_ENGINE"
	AdasyncrequestsetrequestsStatusesEnumParamProcessByEventProcessor AdasyncrequestsetrequestsStatusesEnumParam = "PROCESS_BY_EVENT_PROCESSOR"
	AdasyncrequestsetrequestsStatusesEnumParamSuccess                 AdasyncrequestsetrequestsStatusesEnumParam = "SUCCESS"
	AdasyncrequestsetrequestsStatusesEnumParamUserCanceled            AdasyncrequestsetrequestsStatusesEnumParam = "USER_CANCELED"
	AdasyncrequestsetrequestsStatusesEnumParamUserCanceledDependency  AdasyncrequestsetrequestsStatusesEnumParam = "USER_CANCELED_DEPENDENCY"
)

func (value AdasyncrequestsetrequestsStatusesEnumParam) String() string {
	return string(value)
}

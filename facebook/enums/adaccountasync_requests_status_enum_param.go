package enums

type AdaccountasyncRequestsStatusEnumParam string

const (
	AdaccountasyncRequestsStatusEnumParamError       AdaccountasyncRequestsStatusEnumParam = "ERROR"
	AdaccountasyncRequestsStatusEnumParamExecuting   AdaccountasyncRequestsStatusEnumParam = "EXECUTING"
	AdaccountasyncRequestsStatusEnumParamFinished    AdaccountasyncRequestsStatusEnumParam = "FINISHED"
	AdaccountasyncRequestsStatusEnumParamInitialized AdaccountasyncRequestsStatusEnumParam = "INITIALIZED"
)

func (value AdaccountasyncRequestsStatusEnumParam) String() string {
	return string(value)
}

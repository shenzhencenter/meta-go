package enums

type AdaccountasyncRequestsTypeEnumParam string

const (
	AdaccountasyncRequestsTypeEnumParamAsyncAdgroupCreation AdaccountasyncRequestsTypeEnumParam = "ASYNC_ADGROUP_CREATION"
	AdaccountasyncRequestsTypeEnumParamBatchAPI             AdaccountasyncRequestsTypeEnumParam = "BATCH_API"
	AdaccountasyncRequestsTypeEnumParamDrafts               AdaccountasyncRequestsTypeEnumParam = "DRAFTS"
)

func (value AdaccountasyncRequestsTypeEnumParam) String() string {
	return string(value)
}

package enums

type PagemessageTemplatesStatusEnumParam string

const (
	PagemessageTemplatesStatusEnumParamApproved        PagemessageTemplatesStatusEnumParam = "APPROVED"
	PagemessageTemplatesStatusEnumParamArchived        PagemessageTemplatesStatusEnumParam = "ARCHIVED"
	PagemessageTemplatesStatusEnumParamDeleted         PagemessageTemplatesStatusEnumParam = "DELETED"
	PagemessageTemplatesStatusEnumParamDisabled        PagemessageTemplatesStatusEnumParam = "DISABLED"
	PagemessageTemplatesStatusEnumParamInAppeal        PagemessageTemplatesStatusEnumParam = "IN_APPEAL"
	PagemessageTemplatesStatusEnumParamLimitExceeded   PagemessageTemplatesStatusEnumParam = "LIMIT_EXCEEDED"
	PagemessageTemplatesStatusEnumParamPaused          PagemessageTemplatesStatusEnumParam = "PAUSED"
	PagemessageTemplatesStatusEnumParamPending         PagemessageTemplatesStatusEnumParam = "PENDING"
	PagemessageTemplatesStatusEnumParamPendingDeletion PagemessageTemplatesStatusEnumParam = "PENDING_DELETION"
	PagemessageTemplatesStatusEnumParamRejected        PagemessageTemplatesStatusEnumParam = "REJECTED"
)

func (value PagemessageTemplatesStatusEnumParam) String() string {
	return string(value)
}

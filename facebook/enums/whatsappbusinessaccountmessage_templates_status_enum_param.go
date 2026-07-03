package enums

type WhatsappbusinessaccountmessageTemplatesStatusEnumParam string

const (
	WhatsappbusinessaccountmessageTemplatesStatusEnumParamApproved        WhatsappbusinessaccountmessageTemplatesStatusEnumParam = "APPROVED"
	WhatsappbusinessaccountmessageTemplatesStatusEnumParamArchived        WhatsappbusinessaccountmessageTemplatesStatusEnumParam = "ARCHIVED"
	WhatsappbusinessaccountmessageTemplatesStatusEnumParamDeleted         WhatsappbusinessaccountmessageTemplatesStatusEnumParam = "DELETED"
	WhatsappbusinessaccountmessageTemplatesStatusEnumParamDisabled        WhatsappbusinessaccountmessageTemplatesStatusEnumParam = "DISABLED"
	WhatsappbusinessaccountmessageTemplatesStatusEnumParamInAppeal        WhatsappbusinessaccountmessageTemplatesStatusEnumParam = "IN_APPEAL"
	WhatsappbusinessaccountmessageTemplatesStatusEnumParamLimitExceeded   WhatsappbusinessaccountmessageTemplatesStatusEnumParam = "LIMIT_EXCEEDED"
	WhatsappbusinessaccountmessageTemplatesStatusEnumParamPaused          WhatsappbusinessaccountmessageTemplatesStatusEnumParam = "PAUSED"
	WhatsappbusinessaccountmessageTemplatesStatusEnumParamPending         WhatsappbusinessaccountmessageTemplatesStatusEnumParam = "PENDING"
	WhatsappbusinessaccountmessageTemplatesStatusEnumParamPendingDeletion WhatsappbusinessaccountmessageTemplatesStatusEnumParam = "PENDING_DELETION"
	WhatsappbusinessaccountmessageTemplatesStatusEnumParamRejected        WhatsappbusinessaccountmessageTemplatesStatusEnumParam = "REJECTED"
)

func (value WhatsappbusinessaccountmessageTemplatesStatusEnumParam) String() string {
	return string(value)
}

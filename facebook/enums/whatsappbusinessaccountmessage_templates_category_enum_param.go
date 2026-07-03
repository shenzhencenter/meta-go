package enums

type WhatsappbusinessaccountmessageTemplatesCategoryEnumParam string

const (
	WhatsappbusinessaccountmessageTemplatesCategoryEnumParamAuthentication WhatsappbusinessaccountmessageTemplatesCategoryEnumParam = "AUTHENTICATION"
	WhatsappbusinessaccountmessageTemplatesCategoryEnumParamMarketing      WhatsappbusinessaccountmessageTemplatesCategoryEnumParam = "MARKETING"
	WhatsappbusinessaccountmessageTemplatesCategoryEnumParamUtility        WhatsappbusinessaccountmessageTemplatesCategoryEnumParam = "UTILITY"
)

func (value WhatsappbusinessaccountmessageTemplatesCategoryEnumParam) String() string {
	return string(value)
}

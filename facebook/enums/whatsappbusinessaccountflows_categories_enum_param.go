package enums

type WhatsappbusinessaccountflowsCategoriesEnumParam string

const (
	WhatsappbusinessaccountflowsCategoriesEnumParamAppointmentBooking WhatsappbusinessaccountflowsCategoriesEnumParam = "APPOINTMENT_BOOKING"
	WhatsappbusinessaccountflowsCategoriesEnumParamContactUs          WhatsappbusinessaccountflowsCategoriesEnumParam = "CONTACT_US"
	WhatsappbusinessaccountflowsCategoriesEnumParamCustomerSupport    WhatsappbusinessaccountflowsCategoriesEnumParam = "CUSTOMER_SUPPORT"
	WhatsappbusinessaccountflowsCategoriesEnumParamLeadGeneration     WhatsappbusinessaccountflowsCategoriesEnumParam = "LEAD_GENERATION"
	WhatsappbusinessaccountflowsCategoriesEnumParamOther              WhatsappbusinessaccountflowsCategoriesEnumParam = "OTHER"
	WhatsappbusinessaccountflowsCategoriesEnumParamShopping           WhatsappbusinessaccountflowsCategoriesEnumParam = "SHOPPING"
	WhatsappbusinessaccountflowsCategoriesEnumParamSignIn             WhatsappbusinessaccountflowsCategoriesEnumParam = "SIGN_IN"
	WhatsappbusinessaccountflowsCategoriesEnumParamSignUp             WhatsappbusinessaccountflowsCategoriesEnumParam = "SIGN_UP"
	WhatsappbusinessaccountflowsCategoriesEnumParamSurvey             WhatsappbusinessaccountflowsCategoriesEnumParam = "SURVEY"
)

func (value WhatsappbusinessaccountflowsCategoriesEnumParam) String() string {
	return string(value)
}

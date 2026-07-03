package enums

type IggraphusermessengerProfileFieldsEnumParam string

const (
	IggraphusermessengerProfileFieldsEnumParamAccountLinkingURL          IggraphusermessengerProfileFieldsEnumParam = "ACCOUNT_LINKING_URL"
	IggraphusermessengerProfileFieldsEnumParamCommands                   IggraphusermessengerProfileFieldsEnumParam = "COMMANDS"
	IggraphusermessengerProfileFieldsEnumParamDescription                IggraphusermessengerProfileFieldsEnumParam = "DESCRIPTION"
	IggraphusermessengerProfileFieldsEnumParamGetStarted                 IggraphusermessengerProfileFieldsEnumParam = "GET_STARTED"
	IggraphusermessengerProfileFieldsEnumParamHomeURL                    IggraphusermessengerProfileFieldsEnumParam = "HOME_URL"
	IggraphusermessengerProfileFieldsEnumParamIceBreakers                IggraphusermessengerProfileFieldsEnumParam = "ICE_BREAKERS"
	IggraphusermessengerProfileFieldsEnumParamPersistentMenu             IggraphusermessengerProfileFieldsEnumParam = "PERSISTENT_MENU"
	IggraphusermessengerProfileFieldsEnumParamPlatform                   IggraphusermessengerProfileFieldsEnumParam = "PLATFORM"
	IggraphusermessengerProfileFieldsEnumParamSubjectToNewEuPrivacyRules IggraphusermessengerProfileFieldsEnumParam = "SUBJECT_TO_NEW_EU_PRIVACY_RULES"
	IggraphusermessengerProfileFieldsEnumParamTitle                      IggraphusermessengerProfileFieldsEnumParam = "TITLE"
	IggraphusermessengerProfileFieldsEnumParamWhitelistedDomains         IggraphusermessengerProfileFieldsEnumParam = "WHITELISTED_DOMAINS"
)

func (value IggraphusermessengerProfileFieldsEnumParam) String() string {
	return string(value)
}

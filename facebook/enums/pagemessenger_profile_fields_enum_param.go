package enums

type PagemessengerProfileFieldsEnumParam string

const (
	PagemessengerProfileFieldsEnumParamAccountLinkingURL          PagemessengerProfileFieldsEnumParam = "ACCOUNT_LINKING_URL"
	PagemessengerProfileFieldsEnumParamCommands                   PagemessengerProfileFieldsEnumParam = "COMMANDS"
	PagemessengerProfileFieldsEnumParamDescription                PagemessengerProfileFieldsEnumParam = "DESCRIPTION"
	PagemessengerProfileFieldsEnumParamGetStarted                 PagemessengerProfileFieldsEnumParam = "GET_STARTED"
	PagemessengerProfileFieldsEnumParamHomeURL                    PagemessengerProfileFieldsEnumParam = "HOME_URL"
	PagemessengerProfileFieldsEnumParamIceBreakers                PagemessengerProfileFieldsEnumParam = "ICE_BREAKERS"
	PagemessengerProfileFieldsEnumParamPersistentMenu             PagemessengerProfileFieldsEnumParam = "PERSISTENT_MENU"
	PagemessengerProfileFieldsEnumParamPlatform                   PagemessengerProfileFieldsEnumParam = "PLATFORM"
	PagemessengerProfileFieldsEnumParamSubjectToNewEuPrivacyRules PagemessengerProfileFieldsEnumParam = "SUBJECT_TO_NEW_EU_PRIVACY_RULES"
	PagemessengerProfileFieldsEnumParamTitle                      PagemessengerProfileFieldsEnumParam = "TITLE"
	PagemessengerProfileFieldsEnumParamWhitelistedDomains         PagemessengerProfileFieldsEnumParam = "WHITELISTED_DOMAINS"
)

func (value PagemessengerProfileFieldsEnumParam) String() string {
	return string(value)
}

package enums

type PagecustomUserSettingsParamsEnumParam string

const (
	PagecustomUserSettingsParamsEnumParamPersistentMenu PagecustomUserSettingsParamsEnumParam = "PERSISTENT_MENU"
)

func (value PagecustomUserSettingsParamsEnumParam) String() string {
	return string(value)
}

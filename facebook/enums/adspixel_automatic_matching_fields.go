package enums

type AdspixelAutomaticMatchingFields string

const (
	AdspixelAutomaticMatchingFieldsCountry    AdspixelAutomaticMatchingFields = "country"
	AdspixelAutomaticMatchingFieldsCt         AdspixelAutomaticMatchingFields = "ct"
	AdspixelAutomaticMatchingFieldsDb         AdspixelAutomaticMatchingFields = "db"
	AdspixelAutomaticMatchingFieldsEm         AdspixelAutomaticMatchingFields = "em"
	AdspixelAutomaticMatchingFieldsExternalID AdspixelAutomaticMatchingFields = "external_id"
	AdspixelAutomaticMatchingFieldsFn         AdspixelAutomaticMatchingFields = "fn"
	AdspixelAutomaticMatchingFieldsGe         AdspixelAutomaticMatchingFields = "ge"
	AdspixelAutomaticMatchingFieldsLn         AdspixelAutomaticMatchingFields = "ln"
	AdspixelAutomaticMatchingFieldsPh         AdspixelAutomaticMatchingFields = "ph"
	AdspixelAutomaticMatchingFieldsSt         AdspixelAutomaticMatchingFields = "st"
	AdspixelAutomaticMatchingFieldsZp         AdspixelAutomaticMatchingFields = "zp"
)

func (value AdspixelAutomaticMatchingFields) String() string {
	return string(value)
}

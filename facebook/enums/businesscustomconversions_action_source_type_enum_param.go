package enums

type BusinesscustomconversionsActionSourceTypeEnumParam string

const (
	BusinesscustomconversionsActionSourceTypeEnumParamApp               BusinesscustomconversionsActionSourceTypeEnumParam = "app"
	BusinesscustomconversionsActionSourceTypeEnumParamBusinessMessaging BusinesscustomconversionsActionSourceTypeEnumParam = "business_messaging"
	BusinesscustomconversionsActionSourceTypeEnumParamChat              BusinesscustomconversionsActionSourceTypeEnumParam = "chat"
	BusinesscustomconversionsActionSourceTypeEnumParamEmail             BusinesscustomconversionsActionSourceTypeEnumParam = "email"
	BusinesscustomconversionsActionSourceTypeEnumParamOther             BusinesscustomconversionsActionSourceTypeEnumParam = "other"
	BusinesscustomconversionsActionSourceTypeEnumParamPhoneCall         BusinesscustomconversionsActionSourceTypeEnumParam = "phone_call"
	BusinesscustomconversionsActionSourceTypeEnumParamPhysicalStore     BusinesscustomconversionsActionSourceTypeEnumParam = "physical_store"
	BusinesscustomconversionsActionSourceTypeEnumParamSystemGenerated   BusinesscustomconversionsActionSourceTypeEnumParam = "system_generated"
	BusinesscustomconversionsActionSourceTypeEnumParamWebsite           BusinesscustomconversionsActionSourceTypeEnumParam = "website"
)

func (value BusinesscustomconversionsActionSourceTypeEnumParam) String() string {
	return string(value)
}

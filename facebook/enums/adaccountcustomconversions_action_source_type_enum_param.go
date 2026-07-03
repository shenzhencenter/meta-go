package enums

type AdaccountcustomconversionsActionSourceTypeEnumParam string

const (
	AdaccountcustomconversionsActionSourceTypeEnumParamApp               AdaccountcustomconversionsActionSourceTypeEnumParam = "app"
	AdaccountcustomconversionsActionSourceTypeEnumParamBusinessMessaging AdaccountcustomconversionsActionSourceTypeEnumParam = "business_messaging"
	AdaccountcustomconversionsActionSourceTypeEnumParamChat              AdaccountcustomconversionsActionSourceTypeEnumParam = "chat"
	AdaccountcustomconversionsActionSourceTypeEnumParamEmail             AdaccountcustomconversionsActionSourceTypeEnumParam = "email"
	AdaccountcustomconversionsActionSourceTypeEnumParamOther             AdaccountcustomconversionsActionSourceTypeEnumParam = "other"
	AdaccountcustomconversionsActionSourceTypeEnumParamPhoneCall         AdaccountcustomconversionsActionSourceTypeEnumParam = "phone_call"
	AdaccountcustomconversionsActionSourceTypeEnumParamPhysicalStore     AdaccountcustomconversionsActionSourceTypeEnumParam = "physical_store"
	AdaccountcustomconversionsActionSourceTypeEnumParamSystemGenerated   AdaccountcustomconversionsActionSourceTypeEnumParam = "system_generated"
	AdaccountcustomconversionsActionSourceTypeEnumParamWebsite           AdaccountcustomconversionsActionSourceTypeEnumParam = "website"
)

func (value AdaccountcustomconversionsActionSourceTypeEnumParam) String() string {
	return string(value)
}

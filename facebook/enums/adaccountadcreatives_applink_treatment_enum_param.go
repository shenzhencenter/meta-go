package enums

type AdaccountadcreativesApplinkTreatmentEnumParam string

const (
	AdaccountadcreativesApplinkTreatmentEnumParamAutomatic                    AdaccountadcreativesApplinkTreatmentEnumParam = "automatic"
	AdaccountadcreativesApplinkTreatmentEnumParamDeeplinkWithAppstoreFallback AdaccountadcreativesApplinkTreatmentEnumParam = "deeplink_with_appstore_fallback"
	AdaccountadcreativesApplinkTreatmentEnumParamDeeplinkWithWebFallback      AdaccountadcreativesApplinkTreatmentEnumParam = "deeplink_with_web_fallback"
	AdaccountadcreativesApplinkTreatmentEnumParamWebOnly                      AdaccountadcreativesApplinkTreatmentEnumParam = "web_only"
)

func (value AdaccountadcreativesApplinkTreatmentEnumParam) String() string {
	return string(value)
}

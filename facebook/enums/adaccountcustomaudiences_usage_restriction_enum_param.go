package enums

type AdaccountcustomaudiencesUsageRestrictionEnumParam string

const (
	AdaccountcustomaudiencesUsageRestrictionEnumParamExclusionOnly AdaccountcustomaudiencesUsageRestrictionEnumParam = "EXCLUSION_ONLY"
	AdaccountcustomaudiencesUsageRestrictionEnumParamNone          AdaccountcustomaudiencesUsageRestrictionEnumParam = "NONE"
)

func (value AdaccountcustomaudiencesUsageRestrictionEnumParam) String() string {
	return string(value)
}

package enums

type AdaccounttargetingsuggestionsLimitTypeEnumParam string

const (
	AdaccounttargetingsuggestionsLimitTypeEnumParamBehaviors            AdaccounttargetingsuggestionsLimitTypeEnumParam = "behaviors"
	AdaccounttargetingsuggestionsLimitTypeEnumParamCollegeYears         AdaccounttargetingsuggestionsLimitTypeEnumParam = "college_years"
	AdaccounttargetingsuggestionsLimitTypeEnumParamEducationMajors      AdaccounttargetingsuggestionsLimitTypeEnumParam = "education_majors"
	AdaccounttargetingsuggestionsLimitTypeEnumParamEducationSchools     AdaccounttargetingsuggestionsLimitTypeEnumParam = "education_schools"
	AdaccounttargetingsuggestionsLimitTypeEnumParamEducationStatuses    AdaccounttargetingsuggestionsLimitTypeEnumParam = "education_statuses"
	AdaccounttargetingsuggestionsLimitTypeEnumParamFamilyStatuses       AdaccounttargetingsuggestionsLimitTypeEnumParam = "family_statuses"
	AdaccounttargetingsuggestionsLimitTypeEnumParamHomeValue            AdaccounttargetingsuggestionsLimitTypeEnumParam = "home_value"
	AdaccounttargetingsuggestionsLimitTypeEnumParamIncome               AdaccounttargetingsuggestionsLimitTypeEnumParam = "income"
	AdaccounttargetingsuggestionsLimitTypeEnumParamIndustries           AdaccounttargetingsuggestionsLimitTypeEnumParam = "industries"
	AdaccounttargetingsuggestionsLimitTypeEnumParamInterestedIn         AdaccounttargetingsuggestionsLimitTypeEnumParam = "interested_in"
	AdaccounttargetingsuggestionsLimitTypeEnumParamInterests            AdaccounttargetingsuggestionsLimitTypeEnumParam = "interests"
	AdaccounttargetingsuggestionsLimitTypeEnumParamLifeEvents           AdaccounttargetingsuggestionsLimitTypeEnumParam = "life_events"
	AdaccounttargetingsuggestionsLimitTypeEnumParamLocationCategories   AdaccounttargetingsuggestionsLimitTypeEnumParam = "location_categories"
	AdaccounttargetingsuggestionsLimitTypeEnumParamRelationshipStatuses AdaccounttargetingsuggestionsLimitTypeEnumParam = "relationship_statuses"
	AdaccounttargetingsuggestionsLimitTypeEnumParamUserAdclusters       AdaccounttargetingsuggestionsLimitTypeEnumParam = "user_adclusters"
	AdaccounttargetingsuggestionsLimitTypeEnumParamWorkEmployers        AdaccounttargetingsuggestionsLimitTypeEnumParam = "work_employers"
	AdaccounttargetingsuggestionsLimitTypeEnumParamWorkPositions        AdaccounttargetingsuggestionsLimitTypeEnumParam = "work_positions"
)

func (value AdaccounttargetingsuggestionsLimitTypeEnumParam) String() string {
	return string(value)
}

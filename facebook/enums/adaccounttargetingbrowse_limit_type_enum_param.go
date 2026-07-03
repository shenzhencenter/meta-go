package enums

type AdaccounttargetingbrowseLimitTypeEnumParam string

const (
	AdaccounttargetingbrowseLimitTypeEnumParamBehaviors            AdaccounttargetingbrowseLimitTypeEnumParam = "behaviors"
	AdaccounttargetingbrowseLimitTypeEnumParamCollegeYears         AdaccounttargetingbrowseLimitTypeEnumParam = "college_years"
	AdaccounttargetingbrowseLimitTypeEnumParamEducationMajors      AdaccounttargetingbrowseLimitTypeEnumParam = "education_majors"
	AdaccounttargetingbrowseLimitTypeEnumParamEducationSchools     AdaccounttargetingbrowseLimitTypeEnumParam = "education_schools"
	AdaccounttargetingbrowseLimitTypeEnumParamEducationStatuses    AdaccounttargetingbrowseLimitTypeEnumParam = "education_statuses"
	AdaccounttargetingbrowseLimitTypeEnumParamEthnicAffinity       AdaccounttargetingbrowseLimitTypeEnumParam = "ethnic_affinity"
	AdaccounttargetingbrowseLimitTypeEnumParamFamilyStatuses       AdaccounttargetingbrowseLimitTypeEnumParam = "family_statuses"
	AdaccounttargetingbrowseLimitTypeEnumParamGeneration           AdaccounttargetingbrowseLimitTypeEnumParam = "generation"
	AdaccounttargetingbrowseLimitTypeEnumParamHomeOwnership        AdaccounttargetingbrowseLimitTypeEnumParam = "home_ownership"
	AdaccounttargetingbrowseLimitTypeEnumParamHomeType             AdaccounttargetingbrowseLimitTypeEnumParam = "home_type"
	AdaccounttargetingbrowseLimitTypeEnumParamHomeValue            AdaccounttargetingbrowseLimitTypeEnumParam = "home_value"
	AdaccounttargetingbrowseLimitTypeEnumParamHouseholdComposition AdaccounttargetingbrowseLimitTypeEnumParam = "household_composition"
	AdaccounttargetingbrowseLimitTypeEnumParamIncome               AdaccounttargetingbrowseLimitTypeEnumParam = "income"
	AdaccounttargetingbrowseLimitTypeEnumParamIndustries           AdaccounttargetingbrowseLimitTypeEnumParam = "industries"
	AdaccounttargetingbrowseLimitTypeEnumParamInterestedIn         AdaccounttargetingbrowseLimitTypeEnumParam = "interested_in"
	AdaccounttargetingbrowseLimitTypeEnumParamInterests            AdaccounttargetingbrowseLimitTypeEnumParam = "interests"
	AdaccounttargetingbrowseLimitTypeEnumParamLifeEvents           AdaccounttargetingbrowseLimitTypeEnumParam = "life_events"
	AdaccounttargetingbrowseLimitTypeEnumParamLocationCategories   AdaccounttargetingbrowseLimitTypeEnumParam = "location_categories"
	AdaccounttargetingbrowseLimitTypeEnumParamMoms                 AdaccounttargetingbrowseLimitTypeEnumParam = "moms"
	AdaccounttargetingbrowseLimitTypeEnumParamNetWorth             AdaccounttargetingbrowseLimitTypeEnumParam = "net_worth"
	AdaccounttargetingbrowseLimitTypeEnumParamOfficeType           AdaccounttargetingbrowseLimitTypeEnumParam = "office_type"
	AdaccounttargetingbrowseLimitTypeEnumParamPolitics             AdaccounttargetingbrowseLimitTypeEnumParam = "politics"
	AdaccounttargetingbrowseLimitTypeEnumParamRelationshipStatuses AdaccounttargetingbrowseLimitTypeEnumParam = "relationship_statuses"
	AdaccounttargetingbrowseLimitTypeEnumParamUserAdclusters       AdaccounttargetingbrowseLimitTypeEnumParam = "user_adclusters"
	AdaccounttargetingbrowseLimitTypeEnumParamWorkEmployers        AdaccounttargetingbrowseLimitTypeEnumParam = "work_employers"
	AdaccounttargetingbrowseLimitTypeEnumParamWorkPositions        AdaccounttargetingbrowseLimitTypeEnumParam = "work_positions"
)

func (value AdaccounttargetingbrowseLimitTypeEnumParam) String() string {
	return string(value)
}

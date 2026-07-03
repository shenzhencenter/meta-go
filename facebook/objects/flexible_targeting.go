package objects

type FlexibleTargeting struct {
	Behaviors            *[]IDName `json:"behaviors,omitempty"`
	CollegeYears         *[]uint64 `json:"college_years,omitempty"`
	Connections          *[]IDName `json:"connections,omitempty"`
	CustomAudiences      *[]IDName `json:"custom_audiences,omitempty"`
	EducationMajors      *[]IDName `json:"education_majors,omitempty"`
	EducationSchools     *[]IDName `json:"education_schools,omitempty"`
	EducationStatuses    *[]uint64 `json:"education_statuses,omitempty"`
	EthnicAffinity       *[]IDName `json:"ethnic_affinity,omitempty"`
	FamilyStatuses       *[]IDName `json:"family_statuses,omitempty"`
	FriendsOfConnections *[]IDName `json:"friends_of_connections,omitempty"`
	Generation           *[]IDName `json:"generation,omitempty"`
	HomeOwnership        *[]IDName `json:"home_ownership,omitempty"`
	HomeType             *[]IDName `json:"home_type,omitempty"`
	HomeValue            *[]IDName `json:"home_value,omitempty"`
	HouseholdComposition *[]IDName `json:"household_composition,omitempty"`
	Income               *[]IDName `json:"income,omitempty"`
	Industries           *[]IDName `json:"industries,omitempty"`
	InterestedIn         *[]uint64 `json:"interested_in,omitempty"`
	Interests            *[]IDName `json:"interests,omitempty"`
	LifeEvents           *[]IDName `json:"life_events,omitempty"`
	Moms                 *[]IDName `json:"moms,omitempty"`
	NetWorth             *[]IDName `json:"net_worth,omitempty"`
	OfficeType           *[]IDName `json:"office_type,omitempty"`
	Politics             *[]IDName `json:"politics,omitempty"`
	RelationshipStatuses *[]uint64 `json:"relationship_statuses,omitempty"`
	UserAdclusters       *[]IDName `json:"user_adclusters,omitempty"`
	WorkEmployers        *[]IDName `json:"work_employers,omitempty"`
	WorkPositions        *[]IDName `json:"work_positions,omitempty"`
}

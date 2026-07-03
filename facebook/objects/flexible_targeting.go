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

var FlexibleTargetingFields = struct {
	Behaviors            string
	CollegeYears         string
	Connections          string
	CustomAudiences      string
	EducationMajors      string
	EducationSchools     string
	EducationStatuses    string
	EthnicAffinity       string
	FamilyStatuses       string
	FriendsOfConnections string
	Generation           string
	HomeOwnership        string
	HomeType             string
	HomeValue            string
	HouseholdComposition string
	Income               string
	Industries           string
	InterestedIn         string
	Interests            string
	LifeEvents           string
	Moms                 string
	NetWorth             string
	OfficeType           string
	Politics             string
	RelationshipStatuses string
	UserAdclusters       string
	WorkEmployers        string
	WorkPositions        string
}{
	Behaviors:            "behaviors",
	CollegeYears:         "college_years",
	Connections:          "connections",
	CustomAudiences:      "custom_audiences",
	EducationMajors:      "education_majors",
	EducationSchools:     "education_schools",
	EducationStatuses:    "education_statuses",
	EthnicAffinity:       "ethnic_affinity",
	FamilyStatuses:       "family_statuses",
	FriendsOfConnections: "friends_of_connections",
	Generation:           "generation",
	HomeOwnership:        "home_ownership",
	HomeType:             "home_type",
	HomeValue:            "home_value",
	HouseholdComposition: "household_composition",
	Income:               "income",
	Industries:           "industries",
	InterestedIn:         "interested_in",
	Interests:            "interests",
	LifeEvents:           "life_events",
	Moms:                 "moms",
	NetWorth:             "net_worth",
	OfficeType:           "office_type",
	Politics:             "politics",
	RelationshipStatuses: "relationship_statuses",
	UserAdclusters:       "user_adclusters",
	WorkEmployers:        "work_employers",
	WorkPositions:        "work_positions",
}

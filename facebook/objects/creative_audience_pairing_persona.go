package objects

type CreativeAudiencePairingPersona struct {
	AgeMax  *int      `json:"age_max,omitempty"`
	AgeMin  *int      `json:"age_min,omitempty"`
	Genders *[]uint64 `json:"genders,omitempty"`
}

var CreativeAudiencePairingPersonaFields = struct {
	AgeMax  string
	AgeMin  string
	Genders string
}{
	AgeMax:  "age_max",
	AgeMin:  "age_min",
	Genders: "genders",
}

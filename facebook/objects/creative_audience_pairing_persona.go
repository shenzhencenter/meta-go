package objects

type CreativeAudiencePairingPersona struct {
	AgeMax  *int      `json:"age_max,omitempty"`
	AgeMin  *int      `json:"age_min,omitempty"`
	Genders *[]uint64 `json:"genders,omitempty"`
}

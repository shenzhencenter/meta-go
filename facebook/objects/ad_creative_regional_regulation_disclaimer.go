package objects

type AdCreativeRegionalRegulationDisclaimer struct {
	AustraliaFinserv   *map[string]interface{} `json:"australia_finserv,omitempty"`
	IndiaFinserv       *map[string]interface{} `json:"india_finserv,omitempty"`
	SingaporeUniversal *map[string]interface{} `json:"singapore_universal,omitempty"`
	TaiwanFinserv      *map[string]interface{} `json:"taiwan_finserv,omitempty"`
	TaiwanUniversal    *map[string]interface{} `json:"taiwan_universal,omitempty"`
}

var AdCreativeRegionalRegulationDisclaimerFields = struct {
	AustraliaFinserv   string
	IndiaFinserv       string
	SingaporeUniversal string
	TaiwanFinserv      string
	TaiwanUniversal    string
}{
	AustraliaFinserv:   "australia_finserv",
	IndiaFinserv:       "india_finserv",
	SingaporeUniversal: "singapore_universal",
	TaiwanFinserv:      "taiwan_finserv",
	TaiwanUniversal:    "taiwan_universal",
}

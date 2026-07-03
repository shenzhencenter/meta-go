package objects

type AdCreativeRegionalRegulationDisclaimer struct {
	AustraliaFinserv   *map[string]interface{} `json:"australia_finserv,omitempty"`
	IndiaFinserv       *map[string]interface{} `json:"india_finserv,omitempty"`
	SingaporeUniversal *map[string]interface{} `json:"singapore_universal,omitempty"`
	TaiwanFinserv      *map[string]interface{} `json:"taiwan_finserv,omitempty"`
	TaiwanUniversal    *map[string]interface{} `json:"taiwan_universal,omitempty"`
}

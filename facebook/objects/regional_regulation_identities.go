package objects

type RegionalRegulationIdentities struct {
	AustraliaFinservBeneficiary   *string `json:"australia_finserv_beneficiary,omitempty"`
	AustraliaFinservPayer         *string `json:"australia_finserv_payer,omitempty"`
	IndiaFinservBeneficiary       *string `json:"india_finserv_beneficiary,omitempty"`
	IndiaFinservPayer             *string `json:"india_finserv_payer,omitempty"`
	SingaporeUniversalBeneficiary *string `json:"singapore_universal_beneficiary,omitempty"`
	SingaporeUniversalPayer       *string `json:"singapore_universal_payer,omitempty"`
	TaiwanFinservBeneficiary      *string `json:"taiwan_finserv_beneficiary,omitempty"`
	TaiwanFinservPayer            *string `json:"taiwan_finserv_payer,omitempty"`
	TaiwanUniversalBeneficiary    *string `json:"taiwan_universal_beneficiary,omitempty"`
	TaiwanUniversalPayer          *string `json:"taiwan_universal_payer,omitempty"`
	UniversalBeneficiary          *string `json:"universal_beneficiary,omitempty"`
	UniversalPayer                *string `json:"universal_payer,omitempty"`
}

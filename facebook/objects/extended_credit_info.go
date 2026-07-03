package objects

type ExtendedCreditInfo struct {
	CreditLeft    *string `json:"credit_left,omitempty"`
	CreditRevoked *bool   `json:"credit_revoked,omitempty"`
	CreditUsed    *string `json:"credit_used,omitempty"`
	UsingBizEc    *string `json:"using_biz_ec,omitempty"`
}

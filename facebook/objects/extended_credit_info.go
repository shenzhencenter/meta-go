package objects

type ExtendedCreditInfo struct {
	CreditLeft    *string `json:"credit_left,omitempty"`
	CreditRevoked *bool   `json:"credit_revoked,omitempty"`
	CreditUsed    *string `json:"credit_used,omitempty"`
	UsingBizEc    *string `json:"using_biz_ec,omitempty"`
}

var ExtendedCreditInfoFields = struct {
	CreditLeft    string
	CreditRevoked string
	CreditUsed    string
	UsingBizEc    string
}{
	CreditLeft:    "credit_left",
	CreditRevoked: "credit_revoked",
	CreditUsed:    "credit_used",
	UsingBizEc:    "using_biz_ec",
}

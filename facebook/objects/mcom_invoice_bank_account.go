package objects

type McomInvoiceBankAccount struct {
	NumPendingVerificationAccounts *int                      `json:"num_pending_verification_accounts,omitempty"`
	NumVerifiedAccounts            *int                      `json:"num_verified_accounts,omitempty"`
	PendingVerificationAccounts    *[]map[string]interface{} `json:"pending_verification_accounts,omitempty"`
	VerifiedAccounts               *[]map[string]interface{} `json:"verified_accounts,omitempty"`
}

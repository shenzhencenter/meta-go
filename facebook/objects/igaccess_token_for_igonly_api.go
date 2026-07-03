package objects

type IGAccessTokenForIGOnlyAPI struct {
	AccessToken *string `json:"access_token,omitempty"`
	ExpiresIn   *int    `json:"expires_in,omitempty"`
	TokenType   *string `json:"token_type,omitempty"`
}

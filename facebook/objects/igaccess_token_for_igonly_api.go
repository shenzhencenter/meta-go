package objects

type IGAccessTokenForIGOnlyAPI struct {
	AccessToken *string `json:"access_token,omitempty"`
	ExpiresIn   *int    `json:"expires_in,omitempty"`
	TokenType   *string `json:"token_type,omitempty"`
}

var IGAccessTokenForIGOnlyAPIFields = struct {
	AccessToken string
	ExpiresIn   string
	TokenType   string
}{
	AccessToken: "access_token",
	ExpiresIn:   "expires_in",
	TokenType:   "token_type",
}

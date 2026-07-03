package objects

type IGRefreshAccessTokenForIGOnlyAPI struct {
	AccessToken *string `json:"access_token,omitempty"`
	ExpiresIn   *int    `json:"expires_in,omitempty"`
	Permissions *string `json:"permissions,omitempty"`
	TokenType   *string `json:"token_type,omitempty"`
}

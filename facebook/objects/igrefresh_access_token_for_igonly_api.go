package objects

type IGRefreshAccessTokenForIGOnlyAPI struct {
	AccessToken *string `json:"access_token,omitempty"`
	ExpiresIn   *int    `json:"expires_in,omitempty"`
	Permissions *string `json:"permissions,omitempty"`
	TokenType   *string `json:"token_type,omitempty"`
}

var IGRefreshAccessTokenForIGOnlyAPIFields = struct {
	AccessToken string
	ExpiresIn   string
	Permissions string
	TokenType   string
}{
	AccessToken: "access_token",
	ExpiresIn:   "expires_in",
	Permissions: "permissions",
	TokenType:   "token_type",
}

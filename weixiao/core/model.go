package core

type AccessTokenResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	Ocode       string `json:"ocode"`
	ExpiresIn   int64  `json:"expires_in"`
}

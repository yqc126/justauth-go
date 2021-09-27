package model

/**
 * 授权所需的token
 *
 * @author yadong.zhang (yadong.zhang0415(a)gmail.com)
 * @since 1.8
 */
type AuthToken struct {
	AccessToken          string `json:"access_token"`
	ExpireIn             int    `json:"expire_in"`
	RefreshToken         string `json:"refresh_token"`
	RefreshTokenExpireIn int    `json:"refresh_token_expire_in"`
	Uid                  string `json:"uid"`
	OpenId               string `json:"open_id"`
	AccessCode           string `json:"access_code"`
	UnionId              string `json:"union_id"`

	/**
	 * Google附带属性
	 */
	Scope     string `json:"scope"`
	TokenType string `json:"token_type"`
	IdToken   string `json:"id_token"`

	/**
	 * 小米附带属性
	 */
	MacAlgorithm string `json:"mac_algorithm"`
	MacKey       string `json:"mac_key"`

	/**
	 * 企业微信附带属性
	 *
	 * @since 1.10.0
	 */
	Code string `json:"code"`

	/**
	 * Twitter附带属性
	 *
	 * @since 1.13.0
	 */
	OauthToken             string `json:"oauth_token"`
	OauthTokenSecret       string `json:"oauth_token_secret"`
	UserId                 string `json:"user_id"`
	ScreenName             string `json:"screen_name"`
	OauthCallbackConfirmed bool   `json:"oauth_callback_confirmed"`
}

// AuthToken builder pattern code
type AuthTokenBuilder struct {
	authToken *AuthToken
}

func NewAuthTokenBuilder() *AuthTokenBuilder {
	authToken := &AuthToken{}
	b := &AuthTokenBuilder{authToken: authToken}
	return b
}

func (b *AuthTokenBuilder) AccessToken(accessToken string) *AuthTokenBuilder {
	b.authToken.AccessToken = accessToken
	return b
}

func (b *AuthTokenBuilder) ExpireIn(expireIn int) *AuthTokenBuilder {
	b.authToken.ExpireIn = expireIn
	return b
}

func (b *AuthTokenBuilder) RefreshToken(refreshToken string) *AuthTokenBuilder {
	b.authToken.RefreshToken = refreshToken
	return b
}

func (b *AuthTokenBuilder) RefreshTokenExpireIn(refreshTokenExpireIn int) *AuthTokenBuilder {
	b.authToken.RefreshTokenExpireIn = refreshTokenExpireIn
	return b
}

func (b *AuthTokenBuilder) Uid(uid string) *AuthTokenBuilder {
	b.authToken.Uid = uid
	return b
}

func (b *AuthTokenBuilder) OpenId(openId string) *AuthTokenBuilder {
	b.authToken.OpenId = openId
	return b
}

func (b *AuthTokenBuilder) AccessCode(accessCode string) *AuthTokenBuilder {
	b.authToken.AccessCode = accessCode
	return b
}

func (b *AuthTokenBuilder) UnionId(unionId string) *AuthTokenBuilder {
	b.authToken.UnionId = unionId
	return b
}

func (b *AuthTokenBuilder) Scope(scope string) *AuthTokenBuilder {
	b.authToken.Scope = scope
	return b
}

func (b *AuthTokenBuilder) TokenType(tokenType string) *AuthTokenBuilder {
	b.authToken.TokenType = tokenType
	return b
}

func (b *AuthTokenBuilder) IdToken(idToken string) *AuthTokenBuilder {
	b.authToken.IdToken = idToken
	return b
}

func (b *AuthTokenBuilder) MacAlgorithm(macAlgorithm string) *AuthTokenBuilder {
	b.authToken.MacAlgorithm = macAlgorithm
	return b
}

func (b *AuthTokenBuilder) MacKey(macKey string) *AuthTokenBuilder {
	b.authToken.MacKey = macKey
	return b
}

func (b *AuthTokenBuilder) Code(code string) *AuthTokenBuilder {
	b.authToken.Code = code
	return b
}

func (b *AuthTokenBuilder) OauthToken(oauthToken string) *AuthTokenBuilder {
	b.authToken.OauthToken = oauthToken
	return b
}

func (b *AuthTokenBuilder) OauthTokenSecret(oauthTokenSecret string) *AuthTokenBuilder {
	b.authToken.OauthTokenSecret = oauthTokenSecret
	return b
}

func (b *AuthTokenBuilder) UserId(userId string) *AuthTokenBuilder {
	b.authToken.UserId = userId
	return b
}

func (b *AuthTokenBuilder) ScreenName(screenName string) *AuthTokenBuilder {
	b.authToken.ScreenName = screenName
	return b
}

func (b *AuthTokenBuilder) OauthCallbackConfirmed(oauthCallbackConfirmed bool) *AuthTokenBuilder {
	b.authToken.OauthCallbackConfirmed = oauthCallbackConfirmed
	return b
}

func (b *AuthTokenBuilder) Build() (*AuthToken, error) {
	return b.authToken, nil
}

package model

/**
 * 授权所需的token
 *
 * @author yadong.zhang (yadong.zhang0415(a)gmail.com)
 * @since 1.8
 */
type AuthToken struct {
	accessToken          string
	expireIn             int
	refreshToken         string
	refreshTokenExpireIn int
	uid                  string
	openId               string
	accessCode           string
	unionId              string

	/**
	 * Google附带属性
	 */
	scope     string
	tokenType string
	idToken   string

	/**
	 * 小米附带属性
	 */
	macAlgorithm string
	macKey       string

	/**
	 * 企业微信附带属性
	 *
	 * @since 1.10.0
	 */
	code string

	/**
	 * Twitter附带属性
	 *
	 * @since 1.13.0
	 */
	oauthToken             string
	oauthTokenSecret       string
	userId                 string
	screenName             string
	oauthCallbackConfirmed bool
}

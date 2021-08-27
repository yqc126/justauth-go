package model

/**
 * 授权回调时的参数类
 *
 * @author yadong.zhang (yadong.zhang0415(a)gmail.com)
 * @since 1.8.0
 */
type AuthCallback struct {

	/**
	 * 访问AuthorizeUrl后回调时带的参数code
	 */
	code string

	/**
	 * 访问AuthorizeUrl后回调时带的参数auth_code，该参数目前只使用于支付宝登录
	 */
	auth_code string

	/**
	 * 访问AuthorizeUrl后回调时带的参数state，用于和请求AuthorizeUrl前的state比较，防止CSRF攻击
	 */
	state string

	/**
	 * 华为授权登录接受code的参数名
	 *
	 * @since 1.10.0
	 */
	authorization_code string

	/**
	 * Twitter回调后返回的oauth_token
	 *
	 * @since 1.13.0
	 */
	oauth_token string

	/**
	 * Twitter回调后返回的oauth_verifier
	 *
	 * @since 1.13.0
	 */
	oauth_verifier string
}

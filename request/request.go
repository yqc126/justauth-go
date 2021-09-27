package request

import "github.com/justauth/justauth-go/model"

//AuthRequest
/**
 * JustAuth {@code Request}公共接口，所有平台的{@code Request}都需要实现该接口
 * <p>
 * {@link AuthRequest#authorize()}
 * {@link AuthRequest#authorize(String)}
 * {@link AuthRequest#login(AuthCallback)}
 * {@link AuthRequest#revoke(AuthToken)}
 * {@link AuthRequest#refresh(AuthToken)}
 *
 * @author yadong.zhang (yadong.zhang0415(a)gmail.com)
 * @since 1.8
 */
type AuthRequest interface {

	//Authorize
	/**
	 * 返回带{@code state}参数的授权url，授权回调时会带上这个{@code state}
	 *
	 * @param state state 验证授权流程的参数，可以防止csrf
	 * @return 返回授权地址
	 */
	Authorize(state string) string

	//Login
	/**
	 * 第三方登录
	 *
	 * @param authCallback 用于接收回调参数的实体
	 * @return 返回登录成功后的用户信息
	 */
	Login(authCallback model.AuthCallback) (*model.AuthResponse, error)

	//Revoke
	/**
	 * 撤销授权
	 *
	 * @param authToken 登录成功后返回的Token信息
	 * @return AuthResponse
	 */
	Revoke(authToken model.AuthToken) (model.AuthResponse, error)

	//Refresh
	/**
	 * 刷新access token （续期）
	 *
	 * @param authToken 登录成功后返回的Token信息
	 * @return AuthResponse
	 */
	Refresh(authToken model.AuthToken) (model.AuthResponse, error)
}

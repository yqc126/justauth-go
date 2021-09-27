package request

import (
	"github.com/justauth/justauth-go/model"
)

/**
 * 默认的request处理类
 *
 * @author yadong.zhang (yadong.zhang0415(a)gmail.com)
 * @author yangkai.shen (https://xkcoding.com)
 * @since 1.0.0
 */
/**
 * Gitee登录
 *
 * @author yadong.zhang (yadong.zhang0415(a)gmail.com)
 * @since 1.0.0
 */
type AuthGiteeRequest struct {
	AuthDefaultRequest
}

func NewAuthGiteeRequest() *AuthGiteeRequest {
	return &AuthGiteeRequest{}
}

func (r *AuthGiteeRequest) GetAccessToken(authCallback model.AuthCallback) (*model.AuthToken, error) {
	token := r.DoPostAuthorizationCode(authCallback.Code)

	b := model.AuthTokenBuilder{}
	return b.
		AccessToken(token.AccessToken).
		RefreshToken(token.RefreshToken).
		Scope(token.Scope).
		TokenType(token.TokenType).
		ExpireIn(token.ExpireIn).
		Build()
}

//GetUserInfo
/**
 * 使用token换取用户信息
 *
 * @param authToken token信息
 * @return 用户信息
 * @see AuthDefaultRequest#getAccessToken(AuthCallback)
 */
func (r *AuthGiteeRequest) GetUserInfo(authToken model.AuthToken) (*model.AuthUser, error) {
	userInfo := r.DoGetUserInfo(authToken)
	return &userInfo, nil
}

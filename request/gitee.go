package request

import (
	"encoding/json"
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
	token, err := r.DoPostAuthorizationCode(authCallback.Code)
	if err != nil {
		return nil, err
	}

	respToken := GiteeRespToken{}
	err = json.Unmarshal([]byte(token), &respToken)
	if err != nil {
		return nil, err
	}

	return model.NewAuthTokenBuilder().
		AccessToken(respToken.AccessToken).
		RefreshToken(respToken.RefreshToken).
		Scope(respToken.Scope).
		TokenType(respToken.TokenType).
		ExpireIn(respToken.ExpiresIn).
		Build()
}

type GiteeRespToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
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

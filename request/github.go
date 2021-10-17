package request

import (
	"errors"
	"github.com/justauth/justauth-go/cache"
	"github.com/justauth/justauth-go/config"
	"github.com/justauth/justauth-go/enum/scope"
	"github.com/justauth/justauth-go/model"
	"github.com/justauth/justauth-go/utils"
	"net/url"
)

type AuthGithubRequest struct {
	AuthDefaultRequest
}

func NewAuthGithubRequest(conf config.AuthConfig) (*AuthGithubRequest, error) {
	request := &AuthGithubRequest{}
	err := request.AuthDefaultRequest.Init(conf, request.GetAccessToken, config.GITHUB, &cache.MemoryCache{})
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (r *AuthGithubRequest) GetAccessToken(authCallback model.AuthCallback) (*model.AuthToken, error) {
	response, err := r.DoPostAuthorizationCode(authCallback.Code)
	if err != nil {
		return nil, err
	}

	res, err := url.ParseQuery(response)
	if err != nil {
		return nil, err
	}

	err = r.checkResponse(res.Has("error"), res.Get("error_description"))
	if err != nil {
		return nil, err
	}

	return model.NewAuthTokenBuilder().
		AccessToken(res.Get("access_token")).
		Scope(res.Get("scope")).
		TokenType(res.Get("token_type")).
		Build()
}

func (r *AuthGithubRequest) checkResponse(hasError bool, errorDescription string) error {
	if hasError {
		return errors.New(errorDescription)
	}
	return nil
}

//GetUserInfo
/**
 * 使用token换取用户信息
 *
 * @param authToken token信息
 * @return 用户信息
 * @see AuthDefaultRequest#getAccessToken(AuthCallback)
 */
func (r *AuthGithubRequest) GetUserInfo(authToken model.AuthToken) (*model.AuthUser, error) {
	userInfo := r.DoGetUserInfo(authToken)
	return &userInfo, nil
}

func (r *AuthGithubRequest) Authorize(state string) (string, error) {
	values := url.Values{}
	values.Add("scope", r.GetScopes(" ", true, scope.GetDefault(scope.GithubScopes)))
	authorize, err := r.AuthDefaultRequest.Authorize(state)
	if err != nil {
		return "", err
	}

	return utils.BuildUrl(authorize, values)
}

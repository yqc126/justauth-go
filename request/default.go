package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/justauth/justauth-go/cache"
	cfg "github.com/justauth/justauth-go/config"
	"github.com/justauth/justauth-go/enum"
	e "github.com/justauth/justauth-go/error"
	"github.com/justauth/justauth-go/model"
	"github.com/justauth/justauth-go/utils"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//AuthDefaultRequest
/**
 * 默认的request处理类
 *
 * @author yadong.zhang (yadong.zhang0415(a)gmail.com)
 * @author yangkai.shen (https://xkcoding.com)
 * @since 1.0.0
 */
type AuthDefaultRequest struct {
	config         cfg.AuthConfig
	source         cfg.AuthSource
	authStateCache cache.AuthStateCache
	getAccessToken AccessTokenFunc
	getUserInfo    UserInfoFunc
}

func (r *AuthDefaultRequest) Init(config cfg.AuthConfig, getAccessToken AccessTokenFunc, source cfg.AuthSource, authStateCache cache.AuthStateCache) error {
	r.config = config
	r.source = source
	r.authStateCache = authStateCache
	r.getAccessToken = getAccessToken

	checker := AuthChecker{}
	if !checker.IsSupportedAuth(config, source) {
		return e.SourceNotSupportedError
	}
	// 校验配置合法性
	err := checker.CheckConfig(config, source)
	if err != nil {
		return err
	}
	return nil
}

//Revoke
/**
 * 撤销授权
 *
 * @param authToken 登录成功后返回的Token信息
 * @return AuthResponse
 */
func (r *AuthDefaultRequest) Revoke(authToken model.AuthToken) (model.AuthResponse, error) {
	panic("implement me")
}

//Refresh
/**
 * 刷新access token （续期）
 *
 * @param authToken 登录成功后返回的Token信息
 * @return AuthResponse
 */
func (r *AuthDefaultRequest) Refresh(authToken model.AuthToken) (model.AuthResponse, error) {
	panic("implement me")
}

//GetAccessToken
/**
 * 获取access token
 *
 * @param authCallback 授权成功后的回调参数
 * @return token
 * @see AuthDefaultRequest#authorize()
 * @see AuthDefaultRequest#authorize(String)
 */
//func (r AuthDefaultRequest) GetAccessToken(authCallback model.AuthCallback) (*model.AuthToken, error) {
//	return model.NewAuthTokenBuilder().OpenId("openId").
//		ExpireIn(1000).
//		IdToken("idToken").
//		Scope("scope").
//		RefreshToken("refreshToken").
//		AccessToken("accessToken").
//		Code("code").
//		Build()
//}

//GetUserInfo
/**
 * 使用token换取用户信息
 *
 * @param authToken token信息
 * @return 用户信息
 * @see AuthDefaultRequest#getAccessToken(AuthCallback)
 */
func (r AuthDefaultRequest) GetUserInfo(authToken model.AuthToken) (*model.AuthUser, error) {
	return model.NewAuthUserBuilder().Username("test").
		Nickname("test").
		Gender(enum.MALE).
		Token(authToken).
		Source(r.source.String()).
		Build()
}

//Login
/**
 * 统一的登录入口。当通过{@link AuthDefaultRequest#authorize(String)}授权成功后，会跳转到调用方的相关回调方法中
 * 方法的入参可以使用{@code AuthCallback}，{@code AuthCallback}类中封装好了OAuth2授权回调所需要的参数
 *
 * @param authCallback 用于接收回调参数的实体
 * @return AuthResponse
 */
func (r AuthDefaultRequest) Login(authCallback model.AuthCallback) (*model.AuthResponse, error) {
	checker := AuthChecker{}
	err := checker.CheckCode(r.source, authCallback)
	if err != nil {
		return nil, err
	}

	accessToken, err := r.getAccessToken(authCallback)
	if err != nil {
		return nil, err
	}

	userInfo, err := r.GetUserInfo(*accessToken)
	if err != nil {
		return nil, err
	}
	fmt.Println(userInfo)
	return nil, err
}

//ResponseError
/**
 * 处理{@link AuthDefaultRequest#login(AuthCallback)} 发生异常的情况，统一响应参数
 *
 * @param e 具体的异常
 * @return AuthResponse
 */
func (r AuthDefaultRequest) ResponseError(err error) model.AuthResponse {
	return model.AuthResponse{}
}

//Authorize
/**
 * 返回带{@code state}参数的授权url，授权回调时会带上这个{@code state}
 *
 * @param state state 验证授权流程的参数，可以防止csrf
 * @return 返回授权地址
 * @since 1.9.3
 */
func (r AuthDefaultRequest) Authorize(state string) (string, error) {
	queryParam := url.Values{}
	queryParam.Add("response_type", "code")
	queryParam.Add("client_id", r.config.ClientId)
	queryParam.Add("redirect_uri", r.config.RedirectUri)
	queryParam.Add("state", r.GetRealState(state))

	return utils.BuildUrl(r.source.Authorize(), queryParam)
}

//AccessTokenUrl
/**
 * 返回获取accessToken的url
 *
 * @param code 授权码
 * @return 返回获取accessToken的url
 */
func (r AuthDefaultRequest) AccessTokenUrl(code string) string {

	v := url.Values{}
	v.Add("code", code)
	v.Add("client_id", r.config.ClientId)
	v.Add("client_secret", r.config.ClientSecret)
	v.Add("grant_type", "authorization_code")
	v.Add("redirect_uri", r.config.RedirectUri)
	accessTokenUrl, _ := utils.BuildUrl(r.source.AccessToken(), v)
	return accessTokenUrl
}

//RefreshTokenUrl
/**
 * 返回获取accessToken的url
 *
 * @param refreshToken refreshToken
 * @return 返回获取accessToken的url
 */
func (r AuthDefaultRequest) RefreshTokenUrl(refreshToken string) string {
	return ""
}

//UserInfoUrl
/**
 * 返回获取userInfo的url
 *
 * @param authToken token
 * @return 返回获取userInfo的url
 */
func (r AuthDefaultRequest) UserInfoUrl(authToken model.AuthToken) string {
	v := url.Values{}
	v.Add("access_token", authToken.AccessToken)
	accessTokenUrl, _ := utils.BuildUrl(r.source.UserInfo(), v)
	return accessTokenUrl
}

//RevokeUrl
/**
 * 返回获取revoke authorization的url
 *
 * @param authToken token
 * @return 返回获取revoke authorization的url
 */
func (r AuthDefaultRequest) RevokeUrl(authToken model.AuthToken) string {
	return ""
}

//GetRealState
/**
 * 获取state，如果为空， 则默认取当前日期的时间戳
 *
 * @param state 原始的state
 * @return 返回不为null的state
 */
func (r AuthDefaultRequest) GetRealState(state string) string {
	if len(state) == 0 {
		state = uuid.New().String()
	}

	//TODO 添加缓存
	return state
}

//DoPostAuthorizationCode
/**
 * 通用的 authorizationCode 协议
 *
 * @param code code码
 * @return Response
 */
func (r AuthDefaultRequest) DoPostAuthorizationCode(code string) (string, error) {
	resp, err := http.Post(r.AccessTokenUrl(code), "application/json", nil)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("status code not 200:" + strconv.Itoa(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}

//DoGetAuthorizationCode
/**
 * 通用的 authorizationCode 协议
 *
 * @param code code码
 * @return Response
 */
func (r AuthDefaultRequest) DoGetAuthorizationCode(code string) string {
	return ""
}

//DoPostUserInfo
/**
 * 通用的 用户信息
 *
 * @param authToken token封装
 * @return Response
 */
func (r AuthDefaultRequest) DoPostUserInfo(authToken model.AuthToken) string {
	return ""
}

//DoGetUserInfo
/**
 * 通用的 用户信息
 *
 * @param authToken token封装
 * @return Response
 */
func (r AuthDefaultRequest) DoGetUserInfo(authToken model.AuthToken) model.AuthUser {
	resp, _ := http.Get(r.UserInfoUrl(authToken))
	body, _ := ioutil.ReadAll(resp.Body)
	token := model.AuthUser{}
	_ = json.Unmarshal(body, &token)
	return token
}

//DoGetRevoke
/**
 * 通用的post形式的取消授权方法
 *
 * @param authToken token封装
 * @return Response
 */
func (r AuthDefaultRequest) DoGetRevoke(authToken model.AuthToken) string {
	return ""
}

//GetScopes
/**
 * 获取以 {@code separator}分割过后的 scope 信息
 *
 * @param separator     多个 {@code scope} 间的分隔符
 * @param encode        是否 encode 编码
 * @param defaultScopes 默认的 scope， 当客户端没有配置 {@code scopes} 时启用
 * @return String
 * @since 1.16.7
 */
func (r AuthDefaultRequest) GetScopes(separator string, encode bool, defaultScopes []string) string {
	scopes := r.config.Scopes
	if len(scopes) == 0 {
		if len(defaultScopes) == 0 {
			return ""
		}
		scopes = defaultScopes
	}

	if len(separator) == 0 {
		// 默认为空格
		separator = " "
	}
	scopeStr := strings.Join(scopes, separator)
	if encode {
		scopeStr = html.EscapeString(scopeStr)
	}
	return scopeStr
}

package request

import (
	"github.com/justauth/justauth-go/enum"
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

//protected AuthToken getAccessToken(AuthCallback authCallback) {
//        String response = doPostAuthorizationCode(authCallback.getCode());
//        JSONObject accessTokenObject = JSONObject.parseObject(response);
//        this.checkResponse(accessTokenObject);
//        return AuthToken.builder()
//            .accessToken(accessTokenObject.getString("access_token"))
//            .refreshToken(accessTokenObject.getString("refresh_token"))
//            .scope(accessTokenObject.getString("scope"))
//            .tokenType(accessTokenObject.getString("token_type"))
//            .expireIn(accessTokenObject.getIntValue("expires_in"))
//            .build();
//    }

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

//protected AuthUser getUserInfo(AuthToken authToken) {
//        String userInfo = doGetUserInfo(authToken);
//        JSONObject object = JSONObject.parseObject(userInfo);
//        this.checkResponse(object);
//        return AuthUser.builder()
//            .rawUserInfo(object)
//            .uuid(object.getString("id"))
//            .username(object.getString("login"))
//            .avatar(object.getString("avatar_url"))
//            .blog(object.getString("blog"))
//            .nickname(object.getString("name"))
//            .company(object.getString("company"))
//            .location(object.getString("address"))
//            .email(object.getString("email"))
//            .remark(object.getString("bio"))
//            .gender(AuthUserGender.UNKNOWN)
//            .token(authToken)
//            .source(source.toString())
//            .build();
//    }

//GetUserInfo
/**
 * 使用token换取用户信息
 *
 * @param authToken token信息
 * @return 用户信息
 * @see AuthDefaultRequest#getAccessToken(AuthCallback)
 */
func (r *AuthGiteeRequest) GetUserInfo(authToken model.AuthToken) (*model.AuthUser, error) {
	b := model.AuthUserBuilder{}
	return b.Username("test").
		Nickname("test").
		Gender(enum.MALE).
		Token(authToken).
		Source(r.source.String()).
		Build()
}

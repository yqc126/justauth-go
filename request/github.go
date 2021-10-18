package request

import (
	"encoding/json"
	"errors"
	"github.com/justauth/justauth-go/cache"
	"github.com/justauth/justauth-go/config"
	"github.com/justauth/justauth-go/enum"
	"github.com/justauth/justauth-go/enum/scope"
	"github.com/justauth/justauth-go/model"
	"github.com/justauth/justauth-go/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type AuthGithubRequest struct {
	AuthDefaultRequest
}

func NewAuthGithubRequest(conf config.AuthConfig) (*AuthGithubRequest, error) {
	request := &AuthGithubRequest{}
	option := AuthOption{
		getAccessToken: request.GetAccessToken,
		getUserInfo:    request.GetUserInfo,
	}

	err := request.AuthDefaultRequest.Init(conf, option, config.GITHUB, &cache.MemoryCache{})
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
	request, err := http.NewRequest(http.MethodGet, r.source.UserInfo(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "token "+authToken.AccessToken)
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("status code not 200:" + strconv.Itoa(response.StatusCode))
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	userInfo := GithubUserInfo{}
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil, err
	}

	return model.NewAuthUserBuilder().
		RawUserInfo(body).
		UUID(strconv.Itoa(userInfo.Id)).
		Username(userInfo.Login).
		Avatar(userInfo.AvatarUrl).
		Blog(userInfo.Blog).
		Nickname(userInfo.Name).
		Company(userInfo.Company).
		Location(userInfo.Location).
		Email(userInfo.Email).
		Remark(userInfo.Bio).
		Gender(enum.UNKNOWN).
		Token(authToken).
		Source(r.source.String()).
		Build()
}

type GithubUserInfo struct {
	Login             string    `json:"login"`
	Id                int       `json:"id"`
	NodeId            string    `json:"node_id"`
	AvatarUrl         string    `json:"avatar_url"`
	GravatarId        string    `json:"gravatar_id"`
	Url               string    `json:"url"`
	HtmlUrl           string    `json:"html_url"`
	FollowersUrl      string    `json:"followers_url"`
	FollowingUrl      string    `json:"following_url"`
	GistsUrl          string    `json:"gists_url"`
	StarredUrl        string    `json:"starred_url"`
	SubscriptionsUrl  string    `json:"subscriptions_url"`
	OrganizationsUrl  string    `json:"organizations_url"`
	ReposUrl          string    `json:"repos_url"`
	EventsUrl         string    `json:"events_url"`
	ReceivedEventsUrl string    `json:"received_events_url"`
	Type              string    `json:"type"`
	SiteAdmin         bool      `json:"site_admin"`
	Name              string    `json:"name"`
	Company           string    `json:"company"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             string    `json:"email"`
	Hireable          string    `json:"hireable"`
	Bio               string    `json:"bio"`
	TwitterUsername   string    `json:"twitter_username"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
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

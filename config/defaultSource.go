package config

/**
 * JustAuth内置的各api需要的url， 用枚举类分平台类型管理
 *
 * @author yadong.zhang (yadong.zhang0415(a)gmail.com)
 * @since 1.0
 */
type DefaultSource struct {
	name        string
	authorize   string
	accessToken string
	userInfo    string
}

func (d DefaultSource) Authorize() string {
	return d.authorize
}

func (d DefaultSource) AccessToken() string {
	return d.accessToken
}

func (d DefaultSource) UserInfo() string {
	return d.userInfo
}

func (d DefaultSource) Revoke() string {
	return ""
}

func (d DefaultSource) Refresh() string {
	return ""
}

func (d DefaultSource) GetName() string {
	return ""
}

func (d DefaultSource) String() string {
	return d.name
}

var (
	GITEE = DefaultSource{
		name:        "gitee",
		authorize:   "https://gitee.com/oauth/authorize",
		accessToken: "https://gitee.com/oauth/token",
		userInfo:    "https://gitee.com/api/v5/user",
	}
)

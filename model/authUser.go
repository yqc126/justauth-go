package model

import (
	"encoding/json"
	"github.com/justauth/justauth-go/enum"
)

//AuthUser
/**
 * 授权成功后的用户信息，根据授权平台的不同，获取的数据完整性也不同
 *
 * @author yadong.zhang (yadong.zhang0415(a)gmail.com)
 * @since 1.8
 */

type AuthUser struct {

	/**
	 * 用户第三方系统的唯一id。在调用方集成该组件时，可以用uuid + source唯一确定一个用户
	 *
	 * @since 1.3.3
	 */
	UUID string `json:"uuid"`
	/**
	 * 用户名
	 */
	Username string `json:"username"`
	/**
	 * 用户昵称
	 */
	Nickname string `json:"nickname"`
	/**
	 * 用户头像
	 */
	Avatar string `json:"avatar"`
	/**
	 * 用户网址
	 */
	Blog string `json:"blog"`
	/**
	 * 所在公司
	 */
	Company string `json:"company"`
	/**
	 * 位置
	 */
	Location string `json:"location"`
	/**
	 * 用户邮箱
	 */
	Email string `json:"email"`
	/**
	 * 用户备注（各平台中的用户个人介绍）
	 */
	Remark string `json:"remark"`
	/**
	 * 性别
	 */
	Gender enum.AuthUserGender `json:"gender"`
	/**
	 * 用户来源
	 */
	Source string `json:"source"`
	/**
	 * 用户授权的token信息
	 */
	Token AuthToken `json:"token"`
	/**
	 * 第三方平台返回的原始用户信息
	 */
	RawUserInfo json.RawMessage `json:"raw_user_info"`
}

// AuthUser builder pattern code
type AuthUserBuilder struct {
	authUser *AuthUser
}

func NewAuthUserBuilder() *AuthUserBuilder {
	authUser := &AuthUser{}
	b := &AuthUserBuilder{authUser: authUser}
	return b
}

func (b *AuthUserBuilder) UUID(uuid string) *AuthUserBuilder {
	b.authUser.UUID = uuid
	return b
}

func (b *AuthUserBuilder) Username(username string) *AuthUserBuilder {
	b.authUser.Username = username
	return b
}

func (b *AuthUserBuilder) Nickname(nickname string) *AuthUserBuilder {
	b.authUser.Nickname = nickname
	return b
}

func (b *AuthUserBuilder) Avatar(avatar string) *AuthUserBuilder {
	b.authUser.Avatar = avatar
	return b
}

func (b *AuthUserBuilder) Blog(blog string) *AuthUserBuilder {
	b.authUser.Blog = blog
	return b
}

func (b *AuthUserBuilder) Company(company string) *AuthUserBuilder {
	b.authUser.Company = company
	return b
}

func (b *AuthUserBuilder) Location(location string) *AuthUserBuilder {
	b.authUser.Location = location
	return b
}

func (b *AuthUserBuilder) Email(email string) *AuthUserBuilder {
	b.authUser.Email = email
	return b
}

func (b *AuthUserBuilder) Remark(remark string) *AuthUserBuilder {
	b.authUser.Remark = remark
	return b
}

func (b *AuthUserBuilder) Gender(gender enum.AuthUserGender) *AuthUserBuilder {
	b.authUser.Gender = gender
	return b
}

func (b *AuthUserBuilder) Source(source string) *AuthUserBuilder {
	b.authUser.Source = source
	return b
}

func (b *AuthUserBuilder) Token(token AuthToken) *AuthUserBuilder {
	b.authUser.Token = token
	return b
}

func (b *AuthUserBuilder) RawUserInfo(rawUserInfo json.RawMessage) *AuthUserBuilder {
	b.authUser.RawUserInfo = rawUserInfo
	return b
}

func (b *AuthUserBuilder) Build() (*AuthUser, error) {
	return b.authUser, nil
}

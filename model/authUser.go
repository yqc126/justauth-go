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
	uuid string
	/**
	 * 用户名
	 */
	username string
	/**
	 * 用户昵称
	 */
	nickname string
	/**
	 * 用户头像
	 */
	avatar string
	/**
	 * 用户网址
	 */
	blog string
	/**
	 * 所在公司
	 */
	company string
	/**
	 * 位置
	 */
	location string
	/**
	 * 用户邮箱
	 */
	email string
	/**
	 * 用户备注（各平台中的用户个人介绍）
	 */
	remark string
	/**
	 * 性别
	 */
	gender enum.AuthUserGender
	/**
	 * 用户来源
	 */
	source string
	/**
	 * 用户授权的token信息
	 */
	token AuthToken
	/**
	 * 第三方平台返回的原始用户信息
	 */
	rawUserInfo json.RawMessage
}

package model

import "github.com/justauth/justauth-go/enums"

/**
 * JustAuth统一授权响应类
 *
 * @author yadong.zhang (yadong.zhang0415(a)gmail.com)
 * @since 1.8
 */
type AuthResponse struct {

	/**
	 * 授权响应状态码
	 */
	Code int

	/**
	 * 授权响应信息
	 */
	Msg string

	/**
	 * 授权响应数据，当且仅当 code = 2000 时返回
	 */
	data interface{}
}

/**
 * 是否请求成功
 *
 * @return true or false
 */
func (r AuthResponse) Ok() bool {
	return r.Code == enums.SUCCESS.Code
}

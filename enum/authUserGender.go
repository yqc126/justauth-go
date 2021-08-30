package enum

import "strings"

/**
 * 用户性别
 *
 * @author yadong.zhang (yadong.zhang0415(a)gmail.com)
 * @since 1.8
 */
type AuthUserGender struct {
	Code string
	Desc string
}

var (
	/**
	 * MALE/FAMALE为正常值，通过{@link AuthUserGender#getRealGender(String)}方法获取真实的性别
	 * UNKNOWN为容错值，部分平台不会返回用户性别，为了方便统一，使用UNKNOWN标记所有未知或不可测的用户性别信息
	 */
	MALE    = AuthUserGender{Code: "1", Desc: "男"}
	FEMALE  = AuthUserGender{Code: "0", Desc: "女"}
	UNKNOWN = AuthUserGender{Code: "-1", Desc: "未知"}
)

/**
 * 获取用户的实际性别，常规网站
 *
 * @param originalGender 用户第三方标注的原始性别
 * @return 用户性别
 */

func (g AuthUserGender) GetRealGender(originalGender string) AuthUserGender {
	if len(originalGender) == 0 || UNKNOWN.Code == originalGender {
		return UNKNOWN
	}

	males := []string{"m", "男", "1", "male"}
	for _, male := range males {
		if strings.Contains(originalGender, male) {
			return MALE
		}
	}

	return FEMALE
}

/**
 * 获取微信平台用户的实际性别，0表示未定义，1表示男性，2表示女性
 *
 * @param originalGender 用户第三方标注的原始性别
 * @return 用户性别
 * @since 1.13.2
 */
func (g AuthUserGender) GetWechatRealGender(originalGender string) AuthUserGender {
	if len(originalGender) == 0 || "0" == originalGender {
		return UNKNOWN
	}
	return g.GetRealGender(originalGender)
}

package scope

/**
 * 各个平台 scope 类的统一接口
 *
 * @author yadong.zhang (yadong.zhang0415(a)gmail.com)
 * @version 1.0.0
 * @since 1.15.7
 */
type AuthScope interface {

	/**
	 * 获取字符串 {@code scope}，对应为各平台实际使用的 {@code scope}
	 *
	 * @return String
	 */
	GetScope() string

	/**
	 * 判断当前 {@code scope} 是否为各平台默认启用的
	 *
	 * @return boolean
	 */
	IsDefault() bool
}

func GetDefault(scopes []AuthScope) []string {
	result := make([]string, 0)
	for _, scope := range scopes {
		if scope.IsDefault() {
			result = append(result, scope.GetScope())
		}
	}

	return result
}

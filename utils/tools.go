package utils

import "github.com/justauth/justauth-go/config"

/**
 * 授权配置类的校验器
 *
 * @author yadong.zhang (yadong.zhang0415(a)gmail.com)
 * @since 1.6.1-beta
 */
type AuthChecker struct{}

/**
 * 是否支持第三方登录
 *
 * @param config config
 * @param source source
 * @return true or false
 * @since 1.6.1-beta
 */
func (a AuthChecker) IsSupportedAuth(config config.AuthConfig, source config.AuthSource) bool {
	isSupported := len(config.ClientId) == 0 && len(config.ClientSecret) == 0
	return isSupported
}

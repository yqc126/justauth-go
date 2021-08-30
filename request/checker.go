package request

import (
	cfg "github.com/justauth/justauth-go/config"
	"github.com/justauth/justauth-go/model"
)

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
func (c *AuthChecker) IsSupportedAuth(config cfg.AuthConfig, source cfg.AuthSource) bool {
	return true
}

/**
 * 检查配置合法性。针对部分平台， 对redirect uri有特定要求。一般来说redirect uri都是http://，而对于facebook平台， redirect uri 必须是https的链接
 *
 * @param config config
 * @param source source
 * @since 1.6.1-beta
 */
func (c *AuthChecker) CheckConfig(config cfg.AuthConfig, source cfg.AuthSource) error {
	return nil
}

/**
 * 校验回调传回的code
 * <p>
 * {@code v1.10.0}版本中改为传入{@code source}和{@code callback}，对于不同平台使用不同参数接受code的情况统一做处理
 *
 * @param source   当前授权平台
 * @param callback 从第三方授权回调回来时传入的参数集合
 * @since 1.8.0
 */
func (c *AuthChecker) CheckCode(source cfg.AuthSource, callback model.AuthCallback) error {
	return nil
}

/**
 * 校验回调传回的{@code state}，为空或者不存在
 * <p>
 * {@code state}不存在的情况只有两种：
 * 1. {@code state}已使用，被正常清除
 * 2. {@code state}为前端伪造，本身就不存在
 *
 * @param state          {@code state}一定不为空
 * @param source         {@code source}当前授权平台
 * @param authStateCache {@code authStateCache} state缓存实现
 */
func (c *AuthChecker) CheckState(state string, source cfg.AuthSource, authStateCache struct{}) error {
	return nil
}

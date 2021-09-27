package request

import "testing"

func TestName(t *testing.T) {
	// 创建授权request
	//AuthRequest authRequest = new AuthGiteeRequest(AuthConfig.builder()
	//        .clientId("clientId")
	//        .clientSecret("clientSecret")
	//        .redirectUri("redirectUri")
	//        .build());
	//// 生成授权页面
	//authRequest.authorize("state");
	//// 授权登录后会返回code（auth_code（仅限支付宝））、state，1.8.0版本后，可以用AuthCallback类作为回调接口的参数
	//// 注：JustAuth默认保存state的时效为3分钟，3分钟内未使用则会自动清除过期的state
	//authRequest.login(callback);

}

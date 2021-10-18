package request

import (
	"fmt"
	"github.com/justauth/justauth-go/config"
	"github.com/justauth/justauth-go/model"
	"testing"
)

func TestAuthGithubRequest_GetAccessToken(t *testing.T) {
	request, err := NewAuthGithubRequest(config.AuthConfig{
		ClientId:     "ClientId",
		ClientSecret: "ClientSecret",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	authorizeUrl, err := request.Authorize("state")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(authorizeUrl)
}

func TestAuthDefaultRequest_Login(t *testing.T) {
	request, err := NewAuthGithubRequest(config.AuthConfig{
		ClientId:     "ClientId",
		ClientSecret: "ClientSecret",
		RedirectUri:  "http://localhost:8080/oauth",
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	callback := model.AuthCallback{Code: "da04382c6c032ef7aa8e"}
	response, err := request.Login(callback)
	if err != nil {
		fmt.Println(err)
		return
	}

	request.GetUserInfo(model.AuthToken{AccessToken: "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"})
	fmt.Println(response.Msg)
}

func TestAuthGithubRequest_GetUserInfo(t *testing.T) {
	request, err := NewAuthGithubRequest(config.AuthConfig{
		ClientId:     "ClientId",
		ClientSecret: "ClientSecret",
		RedirectUri:  "http://localhost:8080/oauth",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	userInfo, err := request.GetUserInfo(model.AuthToken{AccessToken: "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"})
	if err != nil {
		fmt.Println(err)
		return
	}
	userInfo.RawUserInfo.MarshalJSON()
	fmt.Println(userInfo)
}

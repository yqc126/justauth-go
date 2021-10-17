package request

import (
	"fmt"
	"github.com/justauth/justauth-go/config"
	"github.com/justauth/justauth-go/model"
	"testing"
)

func TestAuthGithubRequest_GetAccessToken(t *testing.T) {
	request, err := NewAuthGithubRequest(config.AuthConfig{
		ClientId:     "5053fadd20eabbbf8e2b",
		ClientSecret: "1d3ec46ea11ed103b6b3b6442f4b4cc54e6d41eb",
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
		ClientId:     "5053fadd20eabbbf8e2b",
		ClientSecret: "1d3ec46ea11ed103b6b3b6442f4b4cc54e6d41eb",
		RedirectUri:  "http://localhost:8080/oauth",
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	callback := model.AuthCallback{Code: "fae749babc5cdb0305e1"}
	response, err := request.Login(callback)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response.Msg)
}

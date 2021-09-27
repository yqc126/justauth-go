package utils

import (
	"fmt"
	"net/url"
	"testing"
)

func TestBuildUrl(t *testing.T) {
	// return UrlBuilder.fromBaseUrl(source.accessToken())
	//            .queryParam("code", code)
	//            .queryParam("client_id", config.getClientId())
	//            .queryParam("client_secret", config.getClientSecret())
	//            .queryParam("grant_type", "authorization_code")
	//            .queryParam("redirect_uri", config.getRedirectUri())
	//            .build();
	baseUrl1 := "https://www.google.com/test/abc"
	v := url.Values{}
	v.Add("code", "code")
	v.Add("code", "code")
	v.Add("client_id", "config.getClientId()")
	v.Add("client_secret", "config.getClientSecret()")
	v.Add("grant_type", "authorization_code")
	v.Add("redirect_uri", "config.getRedirectUri()")

	buildUrl, err := BuildUrl(baseUrl1, v)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(buildUrl)

}

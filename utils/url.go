package utils

import (
	"net/url"
	"strings"
)

func BuildUrl(baseUrl string, queries url.Values) (string, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}

	values := u.Query()

	for key, query := range queries {
		for _, v := range query {
			values.Add(key, v)
		}
	}

	queryStr := values.Encode()
	prefix := strings.Join([]string{u.Scheme, u.Host + u.Path}, "://")
	return prefix + "?" + queryStr, nil
}

package utils

import (
	"net/url"
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
	queryStr, err = url.QueryUnescape(queryStr)
	if err != nil {
		return "", err
	}

	u.RawQuery = queryStr
	return u.String(), nil
}

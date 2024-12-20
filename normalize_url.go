package main

import (
	"net/url"
)

func normalizarUrl(s string) (string, error) {
	parsedUrl, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	return parsedUrl.Host + parsedUrl.Path, nil
}

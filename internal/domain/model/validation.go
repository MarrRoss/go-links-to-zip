package model

import (
	"net/url"
	"path"
	"strings"
)

func IsValidURL(link string) (bool, *url.URL) {
	parsedLink, err := url.Parse(link)
	if err != nil || parsedLink.Path == "" {
		return false, nil
	}
	return true, parsedLink
}

func IsValidExtension(link string) bool {
	fileFormats := []string{".pdf", ".jpeg", "jpg"}
	for _, value := range fileFormats {
		if strings.ToLower(path.Ext(link)) == value {
			return true
		}
	}
	return false
}

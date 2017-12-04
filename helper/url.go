package helper

import (
	"os"
)

func Url(url string) string {
	return BaseUrl() + "/" + url
}

func BaseUrl() string {
	port := "8080"
	if val, ok := os.LookupEnv("PORT"); ok && val != "" {
		port = val
	}
	return "http://localhost:" + port
}

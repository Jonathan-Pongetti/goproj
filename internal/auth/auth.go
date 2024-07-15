package auth

import (
	"net/http"
	"errors"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error){
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part auth")
	}
	return vals[1], nil
}
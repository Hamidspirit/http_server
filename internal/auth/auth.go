package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Abstract away authentication
// GetAPIKey extracts API Key from headers
// Example:
// Authorization: APIKey {api_key goes here}
func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] == "APIKey" {
		return "", errors.New("first part of auth header malformed ")
	}

	return vals[1], nil
}

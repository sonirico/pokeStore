package btp

import (
	"fmt"
	"strings"
)

func newResponseError(msg string) error {
	return fmt.Errorf(msg)
}

var MalformedResponse = newResponseError("got malformed response from server")

func ParseStatus(payload string) (StatusCode, bool) {
	withoutSpaces := strings.TrimSpace(payload)
	parts := strings.Split(withoutSpaces, responseDelim)
	if len(parts) < 2 {
		return 0, false
	}
	code, ok := GetStatusCode(strings.TrimSpace(parts[1]))
	return code, ok
}

func ParseBody(payload string) string {
	withoutSpaces := strings.TrimSpace(payload)
	parts := strings.Split(withoutSpaces, responseDelim)
	return strings.TrimSpace(parts[1])
}

func ParseResponse(payload string) (*Response, error) {
	withoutSpaces := strings.TrimSpace(payload)
	parts := strings.Split(withoutSpaces, "\n")
	partsSize := len(parts)
	if partsSize < 1 {
		return nil, MalformedResponse
	}
	code, ok := ParseStatus(parts[0])
	if !ok {
		return nil, MalformedResponse
	}
	message := ""
	if partsSize > 1 {
		message = ParseBody(parts[1])
	}
	return NewResponse(code, message), nil
}

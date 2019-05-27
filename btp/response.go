package btp

import (
	"bytes"
	"fmt"
	"strings"
)

type StatusCode uint16

const (
	Ok          StatusCode = 200
	Created                = 201
	NoContent              = 204
	ClientError            = 400
	NotFound               = 404
	Conflict               = 409
	ServerError            = 500
)

const responseDelim = ":"

var statusCodeRegistry = map[string]StatusCode{
	"200": Ok,
	"201": Created,
	"204": NoContent,
	"400": ClientError,
	"404": NotFound,
	"409": Conflict,
	"500": ServerError,
}

func GetStatusCode(raw string) (StatusCode, bool) {
	code, ok := statusCodeRegistry[raw]
	return code, ok
}

type Response struct {
	code  StatusCode
	body  string
	Error error
}

func NewResponse(code StatusCode, content string) *Response {
	return &Response{code: code, body: content}
}

func (res *Response) IsOk() bool {
	return res.code >= 200 && res.code < 400
}

func (res *Response) isKo() bool {
	return !res.IsOk()
}

func (res *Response) String() string {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("Status%s %d", responseDelim, res.code))
	buf.WriteByte('\n')

	strippedBody := strings.TrimSpace(res.body)

	if len(strippedBody) > 0 {
		buf.WriteString(fmt.Sprintf("Message%s %s", responseDelim, res.body))
		buf.WriteByte('\n')
	}

	buf.WriteByte('\n')

	return buf.String()
}

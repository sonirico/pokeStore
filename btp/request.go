package btp

import (
	"bytes"
)

type Request struct {
	Verb     Verb
	BasketId string
	ItemType string
	Error    *RequestError
}

var BadRequest = &Request{Error: newError(ClientError, "BAD REQUEST")}

func (req *Request) IsValid() bool {
	return req.Error == nil
}

func (req *Request) String() string {
	var output bytes.Buffer

	output.WriteString(string(req.Verb))
	output.WriteString(requestDelim)
	output.WriteString(req.BasketId)
	output.WriteString(requestDelim)

	if len(req.ItemType) > 0 {
		output.WriteString(req.ItemType)
		output.WriteString(requestDelim)
	}

	output.WriteByte('\n')

	return output.String()
}

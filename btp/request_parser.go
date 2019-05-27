package btp

import (
	"fmt"
	"strings"
)

type Verb string

const (
	Add      Verb = "ADD"
	Checkout      = "CHECKOUT"
	Create        = "CREATE"
	Drop          = "DROP"
)

const (
	SyntaxError       = "expected Verb<%s> to have, at least, %d operands. Got %d"
	InvalidBasketName = "basket identifier <%s> is invalid"
	ItemTypeNotFound  = "item type <%s> not found"
)

const requestDelim = ";"

type VerbParser func([]string) *Request

type RequestParser struct {
	requestParserRegistry map[Verb]VerbParser
}

func NewRequestParser() *RequestParser {
	requestParser := &RequestParser{
		requestParserRegistry: make(map[Verb]VerbParser),
	}
	requestParser.register(Add, ParseAddRequest)
	requestParser.register(Create, ParseCreateRequest)
	requestParser.register(Checkout, ParseCheckoutRequest)
	requestParser.register(Drop, ParseDropRequest)
	return requestParser
}

func (rp *RequestParser) register(verb Verb, parser VerbParser) {
	rp.requestParserRegistry[verb] = parser
}

func (rp *RequestParser) Parse(payload string) *Request {
	payload = strings.TrimSpace(payload)
	if len(payload) < 1 {
		return BadRequest
	}
	requestParts := strings.Split(payload, requestDelim)

	if len(requestParts) < 1 {
		return BadRequest
	}

	verb := Verb(strings.TrimSpace(requestParts[0]))
	verbParser, ok := rp.requestParserRegistry[verb]

	if !ok {
		return BadRequest
	}

	return verbParser(requestParts)
}

func ParseItemType(payload string) (string, bool) {
	result := strings.TrimSpace(payload)
	if len(result) < 1 {
		return "", false
	}
	return result, true
}

func ParseBasketId(request string) (string, bool) {
	result := strings.TrimSpace(request)
	if len(result) < 1 {
		return "", false
	}
	return result, true
}

func ParseDropRequest(requestParts []string) *Request {
	request := &Request{Verb: Drop}
	if len(requestParts) < 2 {
		request.Error = newError(ClientError, SyntaxError, Drop, 2, len(requestParts))
		return request
	}
	basketId, ok := ParseBasketId(requestParts[1])
	if !ok {
		request.Error = newError(ClientError, InvalidBasketName, requestParts[1])
		return request
	}
	request.BasketId = basketId
	return request
}

func ParseCheckoutRequest(requestParts []string) *Request {
	request := &Request{Verb: Checkout}
	if len(requestParts) < 2 {
		request.Error = newError(ClientError, SyntaxError, Checkout, 2, len(requestParts))
		return request
	}
	basketId, ok := ParseBasketId(requestParts[1])
	if !ok {
		request.Error = newError(ClientError, InvalidBasketName, requestParts[1])
		return request
	}
	request.BasketId = basketId
	return request
}

func ParseAddRequest(requestParts []string) *Request {
	// ADD;{basket-id};{item-type};{amount};
	request := &Request{Verb: Add}
	if len(requestParts) < 3 {
		request.Error = newError(ClientError, SyntaxError, Add, 3, len(requestParts))
		return request
	}
	basketId, ok := ParseBasketId(requestParts[1])
	if !ok {
		request.Error = newError(ClientError, InvalidBasketName, requestParts[1])
		return request
	}
	request.BasketId = basketId
	itemType, ok := ParseItemType(requestParts[2])
	if !ok {
		request.Error = newError(NotFound, ItemTypeNotFound, requestParts[2])
	}
	request.ItemType = itemType
	return request
}

func ParseCreateRequest(requestParts []string) *Request {
	request := &Request{Verb: Create}
	if len(requestParts) < 2 {
		request.Error = newError(ClientError, SyntaxError, Create, 2, len(requestParts))
		return request
	}
	basketId, ok := ParseBasketId(requestParts[1])
	if !ok {
		request.Error = newError(ClientError, InvalidBasketName, requestParts[1])
		return request
	}
	request.BasketId = basketId
	return request
}

type RequestError struct {
	Code    StatusCode
	Message string
}

func newError(code StatusCode, messageTemplate string, params ...interface{}) *RequestError {
	return &RequestError{
		Code:    code,
		Message: fmt.Sprintf(messageTemplate, params...),
	}
}

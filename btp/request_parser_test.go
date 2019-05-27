package btp

import (
	"testing"
)

type RequestParserTest struct {
	input    string
	expected *Request
}

func newParser() *RequestParser {
	return NewRequestParser()
}

func assertRequestEqual(t *testing.T, expected *Request, actual *Request) bool {
	t.Helper()

	if expected.ItemType != actual.ItemType {
		t.Errorf("expected Request.ItemType to be '%s'. Got '%s'",
			expected.ItemType, actual.ItemType)
		return false
	}
	if expected.BasketId != actual.BasketId {
		t.Errorf("expected Request.BasketId to be '%s'. Got '%s'",
			expected.BasketId, actual.BasketId)
		return false
	}
	if expected.Verb != actual.Verb {
		t.Errorf("expected Request.Verb to be '%s'. Got '%s'",
			expected.Verb, actual.Verb)
		return false
	}
	return true
}

func runRequestParserTests(t *testing.T, tests []RequestParserTest) {
	t.Helper()

	parser := newParser()

	for _, test := range tests {
		actualRequest := parser.Parse(test.input)
		if actualRequest.Error != nil {
			if test.expected == nil {
				continue
			}
			requestError := actualRequest.Error
			t.Fatalf("Request.Parse returned error. Code: %d. Message: %s",
				requestError.Code, requestError.Message)
		}
		ok := assertRequestEqual(t, test.expected, actualRequest)
		if !ok {
			t.Fatalf("request payload made tests to fail: %s", test.input)
		}
	}
}

func TestParseAddRequest(t *testing.T) {
	tests := []RequestParserTest{
		{
			"ADD;basket-name;REPELENTE",
			&Request{
				Verb:     Add,
				BasketId: "basket-name",
				ItemType: "REPELENTE",
			},
		},
		{
			"  ADD  ;  basket with spaces  ;  REPELENTE  ",
			&Request{
				Verb:     Add,
				BasketId: "basket with spaces",
				ItemType: "REPELENTE",
			},
		},
		{
			"ADD;basket-name;REPELENTE;;;;",
			&Request{
				Verb:     Add,
				BasketId: "basket-name",
				ItemType: "REPELENTE",
			},
		},
	}

	runRequestParserTests(t, tests)
}

func TestParseCreateRequest(t *testing.T) {
	tests := []RequestParserTest{
		{
			"CREATE;basket-name;",
			&Request{
				Verb:     Create,
				BasketId: "basket-name",
			},
		},
		{
			"  CREATE  ;  basket with spaces  ;  ",
			&Request{
				Verb:     Create,
				BasketId: "basket with spaces",
			},
		},
		{
			"CREATE;basket-name;;;",
			&Request{
				Verb:     Create,
				BasketId: "basket-name",
			},
		},
	}

	runRequestParserTests(t, tests)
}

func TestParseDropRequest(t *testing.T) {
	tests := []RequestParserTest{
		{
			"DROP;basket-name;",
			&Request{
				Verb:     Drop,
				BasketId: "basket-name",
			},
		},
		{
			"  DROP  ;  basket with spaces  ;  ",
			&Request{
				Verb:     Drop,
				BasketId: "basket with spaces",
			},
		},
		{
			"DROP;basket-name;;;",
			&Request{
				Verb:     Drop,
				BasketId: "basket-name",
			},
		},
	}

	runRequestParserTests(t, tests)
}

func TestParseCheckoutRequest(t *testing.T) {
	tests := []RequestParserTest{
		{
			"CHECKOUT;basket-name;",
			&Request{
				Verb:     Checkout,
				BasketId: "basket-name",
			},
		},
		{
			"  CHECKOUT  ;  basket with spaces  ;  ",
			&Request{
				Verb:     Checkout,
				BasketId: "basket with spaces",
			},
		},
		{
			"CHECKOUT;basket-name;;;",
			&Request{
				Verb:     Checkout,
				BasketId: "basket-name",
			},
		},
	}

	runRequestParserTests(t, tests)
}

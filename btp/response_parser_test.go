package btp

import (
	"testing"
)

type ResponseParserTest struct {
	input    string
	expected *Response
}

func assertResponseEqual(t *testing.T, expected *Response, actual *Response) bool {
	t.Helper()

	if expected.code != actual.code {
		t.Errorf("expected Response.code to be '%d'. Got '%d'",
			expected.code, actual.code)
		return false
	}
	if expected.body != actual.body {
		t.Errorf("expected Response.BasketId to be '%s'. Got '%s'",
			expected.body, actual.body)
		return false
	}
	return true
}

func runResponseParserTests(t *testing.T, tests []ResponseParserTest) {
	t.Helper()

	for _, test := range tests {
		actualResponse, err := ParseResponse(test.input)
		if err != nil {
			if test.expected == nil {
				continue
			}
			t.Fatalf(err.Error())
		}
		ok := assertResponseEqual(t, test.expected, actualResponse)
		if !ok {
			t.Fatalf("request payload made tests to fail: %s", test.input)
		}
	}
}

func TestParseOkResponse(t *testing.T) {
	tests := []ResponseParserTest{
		{
			`Status: 200
			Message: foobar
		
			`,
			&Response{
				code: Ok,
				body: "foobar",
			},
		},
		{
			`Status: 200	`,
			&Response{
				code: Ok,
				body: "",
			},
		},
		{
			`Status: 13
			Message:
		
			`,
			nil,
		},
		{
			`Status: 400
			Message: Bad request
			`,
			&Response{
				code: ClientError,
				body: "Bad request",
			},
		},
		{
			`Status: 404
			Message: Not found`,
			&Response{
				code: NotFound,
				body: "Not found",
			},
		},
	}

	runResponseParserTests(t, tests)
}

func TestParseBodyResponseIsOptional(t *testing.T) {
	tests := []ResponseParserTest{
		{
			`Status: 201`,
			&Response{
				code: Created,
				body: "",
			},
		},
	}
	runResponseParserTests(t, tests)
}

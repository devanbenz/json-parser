package main

import (
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "{\"test\": \"hello\"}"

	tests := []struct {
		expectedType TokenType
		value        string
	}{
		{LEFT_BRACKET, "{"},
		{RIGHT_BRACKET, "}"},
		{LEFT_BRACE, "["},
		{RIGHT_BRACE, "]"},
		{COLON, ":"},
		{QUOTATION, "\""},
		{EOF, "EOF"},
	}

	inp := strings.NewReader(input)
	lexer := New(inp)

	for i, tt := range tests {
		tok := lexer.NextToken()

		if tok
	}

}

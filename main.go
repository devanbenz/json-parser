package main

import "io"

const (
	ILLEGAL       = "ILLEGAL"
	EOF           = "EOF"
	LEFT_BRACKET  = "{"
	RIGHT_BRACKET = "}"
	LEFT_BRACE    = "["
	RIGHT_BRACE   = "]"
	COLON         = ":"
	QUOTATION     = "\""
)

type TokenType string

type Lexer struct {
	r                  io.Reader
	prvToken, nxtToken Token
}

type Token struct {
	value     string
	tokenType TokenType
}

func New(input io.Reader) *Lexer {
	return &Lexer{
		input,
		nil,
		nil,
	}
}

func (t *Lexer) NextToken() *Token {

}

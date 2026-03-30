package lexer

import (
	"Interpreter/token"
	"log"
	"testing"
)

func TestLexer(t *testing.T) {
	input := "+{}();"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "IDENT"},
		{token.INT, "INT"},
		{token.PLUS, "+"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
		{token.ILLEGAL, "ILLEGAL"},
		{token.FUNCTION, "FUNCTION"},
		{token.LET, "LET"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.ASSIGN, "="},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			log.Fatal("invalid token type", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			log.Fatal("invalid token literal", i, tt.expectedLiteral, tok.Literal)
		}

	}

}

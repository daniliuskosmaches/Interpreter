package lexer

import "Interpreter/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.ReadChar()
	return l

}

func (l *Lexer) ReadChar() byte {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	return l.ch

}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = NewToken(token.ASSIGN, l.ch)
	case '+':

		tok = NewToken(token.PLUS, l.ch)
	case '{':
		tok = NewToken(token.LBRACE, l.ch)
	case '}':
		tok = NewToken(token.RBRACE, l.ch)
	case ';':
		tok = NewToken(token.SEMICOLON, l.ch)
	case '(':
		tok = NewToken(token.LPAREN, l.ch)
	case ')':
		tok = NewToken(token.RPAREN, l.ch)
	case ',':
		tok = NewToken(token.COMMA, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		tok.Literal = string(l.ch)
		tok.Type = token.IDENT
	}

	return tok
}

func NewToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

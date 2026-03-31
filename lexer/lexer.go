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
	l.skipWhitespace()
	switch l.ch {
	case '=':
		tok = NewToken(token.ASSIGN, l.ch)
	case '+':
		tok = NewToken(token.PLUS, l.ch)
	case '{':
		tok = NewToken(token.LBRACE, l.ch)
	case '}':
		tok = NewToken(token.RBRACE, l.ch)
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
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if l.isDigit() {
			tok.Type = token.INT
			tok.Literal = l.readNumber()

		} else {
			tok = NewToken(token.ILLEGAL, l.ch)
		}

	}
	l.ReadChar()

	return tok
}

func (l *Lexer) readNumber() string {
	position := l.position
	//reading the numbers
	for l.isDigit() {
		l.ReadChar()
	}

	return l.input[position:l.position]

}

func (l *Lexer) isDigit() bool {
	return '0' <= l.ch && l.ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.ch) {
		l.ReadChar()
	}
	return l.input[pos:l.position]
}
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func NewToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.ReadChar()
	}
}

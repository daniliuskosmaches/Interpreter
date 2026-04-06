package ast

import (
	"Interpreter/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{Token: token.Token{
				Type: token.LET, Literal: "let"},
				Name: &Identifier{Token: token.Token{
					Type: token.IDENT, Literal: "MyVar"},
					Value: "MyVar",
				},

				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anothervar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "let MyVar = anotherVar;" {
		t.Errorf("program.String() was %s, expected \"let MyVar = anotherVar;\"", program.String())
	}

}

package parser

import (
	"Interpreter/ast"
	"Interpreter/lexer"
	"log"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `let five = 5;`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if len(program.Statement) == 0 {
		t.Fatalf("program.Statement is empty")
	}
	if program.Statement[0].TokenLiteral() != "let" {
		t.Fatalf("program.Statement[0] is not let. got=%q", program.Statement[0].TokenLiteral())
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		statement := program.Statement[i]
		if statement.TokenLiteral() != tt.expectedIdentifier {
			t.Fatalf("tests[%d] - identifier wrong. expected=%q, got=%q", i, tt.expectedIdentifier, statement.TokenLiteral())
		}

	}
}
func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		log.Fatalf("statement is not let. got=%q", s.TokenLiteral())
		return false
	}
	letstmt, ok := s.(*ast.LetStatement)
	if !ok {
		log.Fatalf("statement is not let. got=%T", s)
	}
	if letstmt.Name.Value != name {
		log.Fatalf("name is not %q. got=%q", name, letstmt.Name.Value)
	}
	if letstmt.TokenLiteral() != name {
		log.Fatalf("s.Name is not %q. got=%q", name, letstmt.TokenLiteral())
	}

	return true
}

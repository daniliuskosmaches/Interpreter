package parser

import (
	"Interpreter/lexer"
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
func TestReturnStatement(t *testing.T) {

}

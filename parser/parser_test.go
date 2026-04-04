package parser

import (
	"Interpreter/ast"
	"Interpreter/lexer"
	"log"
	"testing"
)

func CheckErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %s", msg)
	}
	t.FailNow()
}

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	CheckErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

	}
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10; 
return 7;`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	CheckErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		log.Fatalf("program.Statements does not contain %d statements got = %d ", 3, len(program.Statements))

	}
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			log.Fatalf("statement is not return. got=%T", stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			log.Fatalf("statement is expecting to be  return. got=%q", stmt.TokenLiteral())
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
	if letstmt.Name.TokenLiteral() != name {
		log.Fatalf("s.Name is not %q. got=%q", name, letstmt.TokenLiteral())
	}

	return true
}

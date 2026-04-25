package parser

import (
	"Interpreter/ast"
	"Interpreter/lexer"
	"fmt"
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

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	CheckErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("literal is not ast.IntegerLiteral, got = %f", stmt.Expression)
	}
	if literal.Value != 5 {
		t.Fatalf("literal.Value is not 5, got = %d", literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Fatalf("token literal is not 5, got = %s", literal.TokenLiteral())
	}

}
func TestParsingPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}
	for _, tt := range prefixTests {
		l := lexer.New(tt.input)

		p := New(l)
		program := p.ParseProgram()
		CheckErrors(t, p)
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ExpressionStatement, got %T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("staement is not PrefixExpression , got %T", stmt.Expression)

		}
		if exp.Operator != tt.operator {
			t.Fatalf("expression operator is not %q, got %q", tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return

		}

	}

}
func testIntegerLiteral(t *testing.T, exp ast.Expression, value int64) bool {
	integer, ok := exp.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("integer is not IntegerLiteral, got %T", exp)
	}
	if integer.Value != value {
		t.Fatalf("integer value is not %d, got %d", value, integer.Value)
		return false

	}
	if integer.TokenLiteral() != fmt.Sprint(value) {
		t.Fatalf("tokenliteral is not %d, got %s", value, integer.TokenLiteral())
	}
	return true
}

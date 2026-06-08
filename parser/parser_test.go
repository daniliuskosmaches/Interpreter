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
		integerValue interface{}
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		{"!true;", "!", true},
		{"!false;", "!", false},
	}
	for _, tt := range prefixTests {
		l := lexer.New(tt.input)

		p := New(l)
		program := p.ParseProgram()
		CheckErrors(t, p)
		if len(program.Statements) != 1 {

			t.Fatalf("program statements is not 1, got %d", len(program.Statements))

		}
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
func TestParsingInfixExpression(t *testing.T) {
	infixTests := []struct {
		input     string
		leftValue interface{}
		operator  string

		rightValue interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}
	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		CheckErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("ERROR program.Statements contains %d statements. expected =%d", len(program.Statements), 1)
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("ERROR program.Statements[0] is not ast.Statement, got %T", program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("ERROR statement is not ast.InfixExpression, got = %T", stmt.Expression)
		}
		if !testLiteralExpression(t, exp.Right, tt.rightValue) {
			return
		}
		if !testLiteralExpression(t, exp.Left, tt.leftValue) {
			return
		}

		if exp.Operator != tt.operator {
			t.Fatalf("ERROR exp.Operator is not equal to %s, got = %s", exp.Operator, tt.operator)

		}

	}

}

func TestOperatorPrecedence(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c", "((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		CheckErrors(t, p)
		actual := program.String()
		if actual != tt.expected {
			t.Fatalf("actual = %s, expected = %s", actual, tt.expected)
		}

		if program.String() != tt.expected {
			t.Fatalf("program.String() = %q, expected %q", program.String(), tt.expected)
		}
		fmt.Printf("infix expression: %q\n", tt.input)
	}
}
func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {

	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp is not ast.Identifier, got %T", exp)
		return false

	}

	if ident.Value != value {

		t.Fatalf("ident.value is not %q, got %q", value, ident.Value)
		return false

	}
	if ident.TokenLiteral() != value {
		t.Fatalf("ident.TokenLiteral is not %q, got %q", value, ident.TokenLiteral())
		return false
	}

	return true
}
func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))

	case int64:
		return testIntegerLiteral(t, exp, v)

	case string:
		return testIdentifier(t, exp, v)

	case bool:
		return testBooleanLiteral(t, exp, v)

	}
	t.Errorf("type of %T not supported", expected)
	return false

}
func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("opExp is not ast.OperatorExpression, got %T", exp)
		return false

	}
	if !testLiteralExpression(t, opExp.Left, left) {
		return false

	}
	if opExp.Operator != operator {
		t.Errorf("operator expression is not %q, got %q", operator, opExp.Operator)
		return false
	}
	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}
	return true
}
func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	boolean, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("exp is not ast.Boolean, got %T", exp)
		return false
	}
	if boolean.Value != value {
		t.Errorf("value is not boolean, got %t", value)
		return false
	}
	if boolean.TokenLiteral() != fmt.Sprint(value) {
		t.Errorf("token literal is not %t, got %s", value, boolean.TokenLiteral())
	}
	return true
}

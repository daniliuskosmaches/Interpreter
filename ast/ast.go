package ast

import (
	"Interpreter/token"
	"bytes"
)

type Node interface {
	TokenLiteral() string
	String() string
}
type Statement interface {
	Node
	StatementNode()
}
type Expression interface {
	Node
	expressionNode()
}
type IntegerLiteral struct {
	token token.Token
	Value int64
}

func (i *IntegerLiteral) TokenLiteral() string {
	return i.token.Literal
}
func (i *IntegerLiteral) String() string {
	return i.token.Literal
}
func (i *IntegerLiteral) expressionNode() {
}

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()

}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}
	return " "
}
func (e *ExpressionStatement) expressionNode() {

}

func (e *ExpressionStatement) StatementNode() {

}
func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

func (r *ReturnStatement) StatementNode() {

}
func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}
func (r *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(r.TokenLiteral() + " ")
	out.WriteString(r.ReturnValue.String())
	out.WriteString(" ; ")
	return out.String()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (i *Identifier) String() string {
	return i.Value
}

func (l *LetStatement) StatementNode() {

}
func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}
func (l *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(l.TokenLiteral() + " ")
	out.WriteString(l.Name.String())
	out.WriteString(" = ")
	if l.Value != nil {
		out.WriteString(l.Value.String())
	}
	out.WriteString(" ; ")
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {

}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}

}

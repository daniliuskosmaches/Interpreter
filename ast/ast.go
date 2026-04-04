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

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value string
}

func (l *LetStatement) StatementNode() {

}
func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
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

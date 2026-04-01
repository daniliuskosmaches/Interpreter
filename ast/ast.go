package ast

import "Interpreter/token"

type Node interface {
	TokenLiteral() string
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
	Statement []Statement
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
	if len(p.Statement) > 0 {
		return p.Statement[0].TokenLiteral()
	} else {
		return ""
	}

}

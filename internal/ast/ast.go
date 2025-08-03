// internal/ast/ast.go
package ast

import "github.com/Jablinelang/Jabline/internal/token"

// Node es la interfaz base para todo nodo del AST
type Node interface {
	TokenLiteral() string
}

// Statement representa una instrucción (ej: var = hola)
type Statement interface {
	Node
	statementNode()
}

// Expression representa una expresión (ej: nombre, 5 + 2)
type Expression interface {
	Node
	expressionNode()
}

// Program es el nodo raíz del AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement representa: var = hola
type LetStatement struct {
	Token token.Token // el token 'var'
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// EchoStatement representa: echo(var)
type EchoStatement struct {
	Token token.Token // el token 'echo'
	Value Expression
}

func (es *EchoStatement) statementNode()       {}
func (es *EchoStatement) TokenLiteral() string { return es.Token.Literal }

// Identifier representa variables o nombres de funciones

// ejemplo: Nombre := "Juan"
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// FunctionStatement representa una función con nombre y cuerpo
type FunctionStatement struct {
	Token      token.Token
	Name       *Identifier
	Parameters []*Identifier
	Body       []Statement
}


func (fs *FunctionStatement) statementNode() {}
func (fs *FunctionStatement) TokenLiteral() string {
	return fs.Token.Literal
}

// CallExpression representa una llamada a función
type CallExpression struct {
	Token    token.Token // El token de llamada, ej: IDENTIFIER
	Function *Identifier
}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

// IfStatement representa un condicional if/else
type IfStatement struct {
	Token       token.Token
	Condition   Expression
	Consequence []Statement
	Alternative []Statement
}

func (is *IfStatement) statementNode() {}
func (is *IfStatement) TokenLiteral() string {
	return is.Token.Literal
}

// StructStatement representa una declaración de struct
type StructStatement struct {
	Token token.Token
	Name  *Identifier
	Types []string
}

func (ss *StructStatement) statementNode() {}
func (ss *StructStatement) TokenLiteral() string {
	return ss.Token.Literal
}

// InfixExpression representa una expresión como a + b
type InfixExpression struct {
	Token    token.Token // El operador
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// StringLiteral representa una cadena
type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}
func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

// IntegerLiteral representa un número
type IntegerLiteral struct {
	Token token.Token
	Value int
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

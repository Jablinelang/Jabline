package parser

import (
	"strconv"

	"github.com/Jablinelang/Jabline/internal/ast"
	"github.com/Jablinelang/Jabline/internal/lexer"
	"github.com/Jablinelang/Jabline/internal/token"
)

type Parser struct {
	l       *lexer.Lexer
	curTok  token.Token
	nextTok token.Token
	errors  []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curTok = p.nextTok
	p.nextTok = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curTok.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curTok.Type {
	case token.VAR:
		return p.parseLetStatement()
	case token.ECHO:
		return p.parseEchoStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curTok}

	p.nextToken()
	if p.curTok.Type != token.IDENT {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}

	p.nextToken()
	if p.curTok.Type != token.ASSIGN {
		return nil
	}

	p.nextToken()
	stmt.Value = p.parseExpression()
	return stmt
}

func (p *Parser) parseEchoStatement() *ast.EchoStatement {
	stmt := &ast.EchoStatement{Token: p.curTok}
	p.nextToken()
	stmt.Value = p.parseExpression()
	return stmt
}

func (p *Parser) parseExpression() ast.Expression {
	switch p.curTok.Type {
	case token.IDENT:
		return &ast.Identifier{Token: p.curTok, Value: p.curTok.Literal}
	case token.STRING:
		return &ast.StringLiteral{Token: p.curTok, Value: p.curTok.Literal}
	case token.INT:
		val, _ := strconv.Atoi(p.curTok.Literal)
		return &ast.IntegerLiteral{Token: p.curTok, Value: val}
	default:
		return nil
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

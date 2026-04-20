package parser

import (
	"github.com/dgrachov/monkey/monkey/ast"
	"github.com/dgrachov/monkey/monkey/token"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO: We're skipping the expressions until we encounter a semicolon

	for !p.curTokenIs(token.Semi) {
		p.nextToken()
	}

	return statement
}

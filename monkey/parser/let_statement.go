package parser

import (
	"github.com/dgrachov/monkey/monkey/ast"
	"github.com/dgrachov/monkey/monkey/token"
)

func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.Identifier) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.Assign) {
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon

	for !p.curTokenIs(token.Semi) {
		p.nextToken()
	}

	return statement
}

package parser

import (
	"github.com/dgrachov/monkey/monkey/ast"
	"github.com/dgrachov/monkey/monkey/token"
)

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{Token: p.curToken}

	statement.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.Semi) {
		p.nextToken()
	}

	return statement
}

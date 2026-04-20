package parser

import (
	"github.com/dgrachov/monkey/monkey/ast"
	"github.com/dgrachov/monkey/monkey/token"
)

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	for p.curToken.Type != token.EOF {
		statement := p.parseStatement()

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		p.nextToken()
	}

	return program
}

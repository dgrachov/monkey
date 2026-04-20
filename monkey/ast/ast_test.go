package ast_test

import (
	"testing"

	"github.com/dgrachov/monkey/monkey/ast"
	"github.com/dgrachov/monkey/monkey/token"
)

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.Let, Literal: "let"},

				Name: &ast.Identifier{
					Token: token.Token{Type: token.Identifier, Literal: "myVar"},
					Value: "myVar",
				},

				Value: &ast.Identifier{
					Token: token.Token{Type: token.Identifier, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	programStr := program.String()

	if programStr != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", programStr)
	}
}

package parser

import "github.com/dgrachov/monkey/monkey/ast"

type (
	prefixParseFunc func() ast.Expression
	infixParseFunc  func(ast.Expression) ast.Expression
)

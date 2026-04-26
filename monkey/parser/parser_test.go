package parser_test

import (
	"fmt"
	"testing"

	"github.com/dgrachov/monkey/monkey/ast"
	"github.com/dgrachov/monkey/monkey/lexer"
	"github.com/dgrachov/monkey/monkey/parser"
)

func TestLetStatement(t *testing.T) {
	input := `let x = 5;
let y = 10;
let foobar = 676767;`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	checkParserErrors(t, p)

	statementCount := len(program.Statements)
	if statementCount != 3 {
		t.Fatalf("expected program.Statements to contain 3 statements. got=%d", statementCount)
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("tests[%d]", i), func(t *testing.T) {
			testLetStatement(t, program.Statements[i], tt.expectedIdentifier)
		})
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) {
	if s.TokenLiteral() != "let" {
		t.Errorf("expected s.TokenLiteral to be 'let'. got=%q", s.TokenLiteral())
		return
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("expected *ast.LetStatement. got=%T", s)
		return
	}

	if letStmt.Name.Value != name {
		t.Errorf("expected s.Name.Value to be %s. got=%s", name, letStmt.Name.Value)
		return
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("expected s.Name to be %s. got=%s", name, letStmt.Name)
		return
	}
}

func TestReturnStatement(t *testing.T) {
	input := `return 5;
return 10;
return 676767;`

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	checkParserErrors(t, p)

	statementCount := len(program.Statements)
	if statementCount != 3 {
		t.Fatalf("expected program.Statements to contain 3 statements. got=%d", statementCount)
	}

	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statement is not *ast.ReturnStatement. got=%T", statement)
			continue
		}

		if returnStatement.TokenLiteral() != "return" {
			t.Errorf("expected returnStatement.TokenLiteral to be 'return'. got=%q", returnStatement.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	checkParserErrors(t, p)

	statementCount := len(program.Statements)
	if statementCount != 1 {
		t.Fatalf("expected program.Statements to contain 1 statement. got=%d", statementCount)
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expected statement to be *ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	identifier, ok := statement.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expected expression to be *ast.Identifier. got=%T", statement.Expression)
	}

	if identifier.Value != "foobar" {
		t.Errorf("expected identifier.Value to be 'foobar'. got=%s", identifier.TokenLiteral())
	}

	if identifier.TokenLiteral() != "foobar" {
		t.Errorf("expected identifier.TokenLiteral to be 'foobar'. got=%s", identifier.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	checkParserErrors(t, p)

	statementCount := len(program.Statements)
	if statementCount != 1 {
		t.Fatalf("expected program.Statements to contain 1 statement. got=%d", statementCount)
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expected statement to be *ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	literal, ok := statement.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("expected expression literal to be *ast.IntegerLiteral. got=%T", statement.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("expected expression literal value to be %d. got=%d", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("expected literal.TokenLiteral to be %s. got=%s", "5", literal.TokenLiteral())
	}
}

func checkParserErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()
	errorCount := len(errors)

	if errorCount == 0 {
		return
	}

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.Errorf("parser has %d errors", errorCount)

	t.FailNow()
}

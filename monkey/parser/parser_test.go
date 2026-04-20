package parser

import (
	"fmt"
	"testing"

	"github.com/dgrachov/monkey/monkey/ast"
	"github.com/dgrachov/monkey/monkey/lexer"
)

func TestLetStatement(t *testing.T) {
	input := `let x = 5;
let y = 10;
let foobar = 676767;`

	l := lexer.New(input)
	p := New(l)

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

func checkParserErrors(t *testing.T, p *Parser) {
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

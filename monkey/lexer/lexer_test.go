package lexer_test

import (
	"fmt"
	"testing"

	"github.com/dgrachov/monkey/monkey/lexer"
	"github.com/dgrachov/monkey/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);

!-/*5;

5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}
	
10 == 10;
10 != 9;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.Let, "let"},
		{token.Identifier, "five"},
		{token.Assign, "="},
		{token.Integer, "5"},
		{token.Semi, ";"},
		{token.Let, "let"},
		{token.Identifier, "ten"},
		{token.Assign, "="},
		{token.Integer, "10"},
		{token.Semi, ";"},
		{token.Let, "let"},
		{token.Identifier, "add"},
		{token.Assign, "="},
		{token.Function, "fn"},
		{token.LParen, "("},
		{token.Identifier, "x"},
		{token.Comma, ","},
		{token.Identifier, "y"},
		{token.RParen, ")"},
		{token.LBrace, "{"},
		{token.Identifier, "x"},
		{token.Plus, "+"},
		{token.Identifier, "y"},
		{token.Semi, ";"},
		{token.RBrace, "}"},
		{token.Semi, ";"},
		{token.Let, "let"},
		{token.Identifier, "result"},
		{token.Assign, "="},
		{token.Identifier, "add"},
		{token.LParen, "("},
		{token.Identifier, "five"},
		{token.Comma, ","},
		{token.Identifier, "ten"},
		{token.RParen, ")"},
		{token.Semi, ";"},
		{token.Bang, "!"},
		{token.Minus, "-"},
		{token.Slash, "/"},
		{token.Asterisk, "*"},
		{token.Integer, "5"},
		{token.Semi, ";"},
		{token.Integer, "5"},
		{token.LessThan, "<"},
		{token.Integer, "10"},
		{token.GreaterThan, ">"},
		{token.Integer, "5"},
		{token.Semi, ";"},
		{token.If, "if"},
		{token.LParen, "("},
		{token.Integer, "5"},
		{token.LessThan, "<"},
		{token.Integer, "10"},
		{token.RParen, ")"},
		{token.LBrace, "{"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.Semi, ";"},
		{token.RBrace, "}"},
		{token.Else, "else"},
		{token.LBrace, "{"},
		{token.Return, "return"},
		{token.False, "false"},
		{token.Semi, ";"},
		{token.RBrace, "}"},
		{token.Integer, "10"},
		{token.Eq, "=="},
		{token.Integer, "10"},
		{token.Semi, ";"},
		{token.Integer, "10"},
		{token.NotEq, "!="},
		{token.Integer, "9"},
		{token.Semi, ";"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		t.Run(fmt.Sprintf("tests[%d]", i), func(t *testing.T) {
			tok := l.NextToken()
			if tok.Type != tt.expectedType {
				t.Errorf("tokentype wrong. expected=%q, got=%q", tt.expectedType, tok.Type)
			}

			if tok.Literal != tt.expectedLiteral {
				t.Errorf("literal wrong. expected=%q, got=%q", tt.expectedLiteral, tok.Literal)
			}
		})
	}
}

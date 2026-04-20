package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals

	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	// Operators

	ASSIGN      = "="
	PLUS        = "+"
	MINUS       = "-"
	BANG        = "!"
	ASTERISK    = "*"
	SLASH       = "/"
	LESSTHAN    = "<"
	GREATERTHAN = ">"

	// Delimiters

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

package token

type TokenType uint

const (
	Illegal TokenType = iota
	EOF

	Identifier
	Integer
	True
	False
	If
	Else
	Return

	Assign
	Plus
	Minus
	Bang
	Asterisk
	Slash
	LessThan
	GreaterThan
	Eq
	NotEq

	Comma
	Semi
	LParen
	RParen
	LBrace
	RBrace

	Function
	Let
)

var tokenNames = map[TokenType]string{
	Illegal: "ILLEGAL",
	EOF:     "EOF",

	Identifier: "IDENTIFIER",
	Integer:    "INTEGER",
	True:       "TRUE",
	False:      "FALSE",
	If:         "IF",
	Else:       "ELSE",
	Return:     "RETURN",

	Assign:      "=",
	Plus:        "+",
	Minus:       "-",
	Bang:        "!",
	Asterisk:    "*",
	Slash:       "/",
	LessThan:    "<",
	GreaterThan: ">",
	Eq:          "==",
	NotEq:       "!=",

	Comma:  ",",
	Semi:   ";",
	LParen: "(",
	RParen: ")",
	LBrace: "{",
	RBrace: "}",

	Function: "FUNCTION",
	Let:      "LET",
}

func (t TokenType) String() string {
	if name, ok := tokenNames[t]; ok {
		return name
	}

	return "UNKNOWN"
}

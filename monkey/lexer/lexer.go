package lexer

import "github.com/dgrachov/monkey/monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}

	l.readChar()

	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.Integer
			return tok
		} else {
			tok = newToken(token.Illegal, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.Eq, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.Assign, l.ch)
		}
	case '+':
		tok = newToken(token.Plus, l.ch)
	case '-':
		tok = newToken(token.Minus, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NotEq, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.Bang, l.ch)
		}
	case '*':
		tok = newToken(token.Asterisk, l.ch)
	case '/':
		tok = newToken(token.Slash, l.ch)
	case '<':
		tok = newToken(token.LessThan, l.ch)
	case '>':
		tok = newToken(token.GreaterThan, l.ch)
	case ',':
		tok = newToken(token.Comma, l.ch)
	case ';':
		tok = newToken(token.Semi, l.ch)
	case '(':
		tok = newToken(token.LParen, l.ch)
	case ')':
		tok = newToken(token.RParen, l.ch)
	case '{':
		tok = newToken(token.LBrace, l.ch)
	case '}':
		tok = newToken(token.RBrace, l.ch)
	}

	l.readChar()

	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

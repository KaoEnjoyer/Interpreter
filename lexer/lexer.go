package lexer

import (
	"Interpreter/token"
	"fmt"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func InitLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.ch = input[0]
	l.position = 0
	l.readPosition = 1

	return l
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

// funkcja pomocnicza tworząca token
func newToken(_tokenType token.TokenType, _charLiteral byte) token.Token {
	return token.Token{Type: _tokenType, Literal: string(_charLiteral)}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
		if l.peekChar() == '=' {
			l.readChar()
			tok.Type = token.EQUAL
			tok.Literal = fmt.Sprintf("%v%v", l.ch, l.input[l.position-1])
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)
		//not
		if l.peekChar() == '=' {
			l.readChar()
			tok.Type = token.NOT
			tok.Literal = string(l.ch + l.input[l.position-1])
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isDigit(l.ch) {
			tok.Literal = l.readNum()
			tok.Type = token.INT
			return tok
		}
		if isLetter(l.ch) {
			tok.Literal = l.readId()
			tok.Type = token.LookUpIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}

	l.readChar()
	return tok
}

func (l *Lexer) readId() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func (l *Lexer) readNum() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

package lexer

import "github.com/135yshr/t_zungo/token"

// Lexer is zungo's lexer
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New is new instance of Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken is read next token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	}

	l.readChar()
	return tok
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

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{tokenType, string(ch)}
}

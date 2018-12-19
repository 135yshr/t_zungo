package lexer

import "github.com/135yshr/t_zungo/token"

// Lexer is zungo's lexer
type Lexer struct {
	input string
}

// New is new instance of Lexer
func New(input string) *Lexer {
	return &Lexer{input}
}

// NextToken is read next token
func (l *Lexer) NextToken() *token.Token {
	return &token.Token{token.ASSIGN, "="}
}

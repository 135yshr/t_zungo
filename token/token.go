package token

const (
	// ILLEGAL is unexpected value
	ILLEGAL = "ILLEGAL"

	// EOF is end of file
	EOF = "EOF"

	// INDENT is indent
	INDENT = ""

	// INT is int value
	INT = "INT"

	// ASSIGN is =
	ASSIGN = "="

	// PLUS is +
	PLUS = "+"

	// COMMA is ,
	COMMA = ","

	// SEMICOLON is ;
	SEMICOLON = ";"

	// LPAREN is (
	LPAREN = "("

	// RPAREN is )
	RPAREN = ")"
)

// TokenType is token's type
type TokenType string

// Token is zungo tokens
type Token struct {
	Type    TokenType
	Literal string
}

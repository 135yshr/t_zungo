package token

const (
	// ILLEGAL is unexpected value
	ILLEGAL = "ILLEGAL"

	// EOF is end of file
	EOF = "EOF"

	// IDENT is indent
	IDENT = ""

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

	// LBRACE is {
	LBRACE = "{"

	// RBRACE is {
	RBRACE = "}"

	// FUNCTION is function
	FUNCTION = "FUNCTION"

	// LET is variable declaration
	LET = "LET"
)

// TokenType is token's type
type TokenType string

// Token is zungo tokens
type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent is refer to identifiers
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

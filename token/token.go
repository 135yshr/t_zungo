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

	// MINUS is -
	MINUS = "-"

	// BANG is !
	BANG = "!"

	// ASTERISK is *
	ASTERISK = "*"

	// SLASH is !
	SLASH = "/"

	// LT is <
	LT = "<"

	// GT is >
	GT = ">"

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

	// TRUE is true declatration
	TRUE = "TRUE"

	// FALSE is false declatration
	FALSE = "FALSE"

	// IF is if declatation
	IF = "IF"

	// ELSE is else declatation
	ELSE = "ELSE"

	// RETURN is return declatation
	RETURN = "RETURN"
)

// TokenType is token's type
type TokenType string

// Token is zungo tokens
type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

// LookupIdent is refer to identifiers
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

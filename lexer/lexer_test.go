package lexer

import (
	"testing"

	"github.com/135yshr/t_zungo/token"
)

func TestNextToken(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []token.Token
	}{
		{
			name: "=+(){},;を渡して解析できること",
			args: args{
				input: `=+(){},;`,
			},
			want: []token.Token{
				{token.ASSIGN, "="},
				{token.PLUS, "+"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
		{
			name: "sample コードを渡して解析できること",
			args: args{
				input: `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
`,
			},
			want: []token.Token{
				{token.LET, "let"},
				{token.IDENT, "five"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "ten"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "add"},
				{token.ASSIGN, "="},
				{token.FUNCTION, "fn"},
				{token.LPAREN, "("},
				{token.IDENT, "x"},
				{token.COMMA, ","},
				{token.IDENT, "y"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.IDENT, "x"},
				{token.PLUS, "+"},
				{token.IDENT, "y"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.SEMICOLON, ";"},
				{token.LET, "let"},
				{token.IDENT, "result"},
				{token.ASSIGN, "="},
				{token.IDENT, "add"},
				{token.LPAREN, "("},
				{token.IDENT, "five"},
				{token.COMMA, ","},
				{token.IDENT, "ten"},
				{token.RPAREN, ")"},
				{token.SEMICOLON, ";"},
				{token.EOF, ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := New(tt.args.input)
			for i, tt := range tt.want {
				tok := sut.NextToken()
				if tok.Type != tt.Type {
					t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q",
						i, tt.Type, tok.Type)
				}
				if tok.Literal != tt.Literal {
					t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q",
						i, tt.Literal, tok.Literal)
				}
			}
		})
	}
}

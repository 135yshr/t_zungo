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
				input: `=`,
			},
			want: []token.Token{
				{token.ASSIGN, "="},
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

package lexer

import (
	"fmt"
	"testing"

	"jay/jay/token"

	"github.com/stretchr/testify/assert"
)

func TestNextToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedTokens []token.Token
	}{
		{
			name:  "one",
			input: `=+(){},;`,
			expectedTokens: []token.Token{
				{
					Type:    token.ASSIGN,
					Literal: "=",
				},
				{
					Type:    token.PLUS,
					Literal: "+",
				},
				{
					Type:    token.LPAREN,
					Literal: "(",
				},
				{
					Type:    token.RPAREN,
					Literal: ")",
				},
				{
					Type:    token.LBRACE,
					Literal: "{",
				},
				{
					Type:    token.RBRACE,
					Literal: "}",
				},
				{
					Type:    token.COMMA,
					Literal: ",",
				},
				{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				{
					Type:    token.EOF,
					Literal: "",
				},
			},
		},
		{
			name: "two",
			input: `let five = 5;
			let ten = 10;

			let add = fn(x, y) {
				x + y;
			};

			let result = add(five, ten);
			!-/*5;
			5 < 10 > 5;

			if (5 < 10) {
				return true;
			} else {
				return false;
			}`,
			expectedTokens: []token.Token{
				{
					Type:    token.LET,
					Literal: "let",
				},
				{
					Type:    token.IDENT,
					Literal: "five",
				},
				{
					Type:    token.ASSIGN,
					Literal: "=",
				},
				{
					Type:    token.INT,
					Literal: "5",
				},
				{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				{
					Type:    token.LET,
					Literal: "let",
				},
				{
					Type:    token.IDENT,
					Literal: "ten",
				},
				{
					Type:    token.ASSIGN,
					Literal: "=",
				},
				{
					Type:    token.INT,
					Literal: "10",
				},
				{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				{
					Type:    token.LET,
					Literal: "let",
				},
				{
					Type:    token.IDENT,
					Literal: "add",
				},
				{
					Type:    token.ASSIGN,
					Literal: "=",
				},
				{
					Type:    token.FUNCTION,
					Literal: "fn",
				},
				{
					Type:    token.LPAREN,
					Literal: "(",
				},
				{
					Type:    token.IDENT,
					Literal: "x",
				},
				{
					Type:    token.COMMA,
					Literal: ",",
				},
				{
					Type:    token.IDENT,
					Literal: "y",
				},
				{
					Type:    token.RPAREN,
					Literal: ")",
				},
				{
					Type:    token.LBRACE,
					Literal: "{",
				},
				{
					Type:    token.IDENT,
					Literal: "x",
				},
				{
					Type:    token.PLUS,
					Literal: "+",
				},
				{
					Type:    token.IDENT,
					Literal: "y",
				},
				{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				{
					Type:    token.RBRACE,
					Literal: "}",
				},
				{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				{
					Type:    token.LET,
					Literal: "let",
				},
				{
					Type:    token.IDENT,
					Literal: "result",
				},
				{
					Type:    token.ASSIGN,
					Literal: "=",
				},
				{
					Type:    token.IDENT,
					Literal: "add",
				},
				{
					Type:    token.LPAREN,
					Literal: "(",
				},
				{
					Type:    token.IDENT,
					Literal: "five",
				},
				{
					Type:    token.COMMA,
					Literal: ",",
				},
				{
					Type:    token.IDENT,
					Literal: "ten",
				},
				{
					Type:    token.RPAREN,
					Literal: ")",
				},
				{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				{
					Type:    token.BANG,
					Literal: "!",
				},
				{
					Type:    token.MINUS,
					Literal: "-",
				},
				{
					Type:    token.SLASH,
					Literal: "/",
				},
				{
					Type:    token.ASTERISK,
					Literal: "*",
				},
				{
					Type:    token.INT,
					Literal: "5",
				},
				{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				{
					Type:    token.INT,
					Literal: "5",
				},
				{
					Type:    token.LT,
					Literal: "<",
				},
				{
					Type:    token.INT,
					Literal: "10",
				},
				{
					Type:    token.GT,
					Literal: ">",
				},
				{
					Type:    token.INT,
					Literal: "5",
				},
				{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				{
					Type:    token.IF,
					Literal: "if",
				},
				{
					Type:    token.LPAREN,
					Literal: "(",
				},
				{
					Type:    token.INT,
					Literal: "5",
				},
				{
					Type:    token.LT,
					Literal: "<",
				},
				{
					Type:    token.INT,
					Literal: "10",
				},
				{
					Type:    token.RPAREN,
					Literal: ")",
				},
				{
					Type:    token.LBRACE,
					Literal: "{",
				},
				{
					Type:    token.RETURN,
					Literal: "return",
				},
				{
					Type:    token.TRUE,
					Literal: "true",
				},
				{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				{
					Type:    token.RBRACE,
					Literal: "}",
				},
				{
					Type:    token.ELSE,
					Literal: "else",
				},
				{
					Type:    token.LBRACE,
					Literal: "{",
				},
				{
					Type:    token.RETURN,
					Literal: "return",
				},
				{
					Type:    token.FALSE,
					Literal: "false",
				},
				{
					Type:    token.SEMICOLON,
					Literal: ";",
				},
				{
					Type:    token.RBRACE,
					Literal: "}",
				},
				{
					Type:    token.EOF,
					Literal: "",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			l := New(tt.input)

			for _, expectedToken := range tt.expectedTokens {
				token := l.NextToken()

				assert.Equal(t, expectedToken.Type, token.Type, fmt.Sprintf("exepcted: %s, got: %s", expectedToken.Type, token.Type))

				assert.Equal(t, expectedToken.Literal, token.Literal, fmt.Sprintf("exepcted: %s, got: %s", expectedToken.Literal, token.Literal))
			}
		})
	}
}

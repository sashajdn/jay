package ast

import (
	"jay/jay/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAST_String(t *testing.T) {
	t.Parallel()

	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{
					Type:    token.LET,
					Literal: "let",
				},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}

	require.Equal(t, `let myVar = anotherVar;`, program.String())
}

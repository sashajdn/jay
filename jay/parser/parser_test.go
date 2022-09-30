package parser

import (
	"fmt"
	"jay/jay/ast"
	"jay/jay/lexer"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLetStatements(t *testing.T) {
	t.Parallel()

	input := `
let x = 5;
let y = 10;
let foobar = 838338;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	require.NotNil(t, program)

	assert.Len(t, program.Statements, 3, "expecting %d statements, got %d", 3, len(program.Statements))

	tests := []struct {
		expectedIdentfier string
	}{
		{
			expectedIdentfier: "x",
		},
		{
			expectedIdentfier: "y",
		},
		{
			expectedIdentfier: "foobar",
		},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		testLetStatements(t, stmt, tt.expectedIdentfier)
	}
}

func testLetStatements(t *testing.T, s ast.Statement, name string) {
	assert.Equal(t, s.TokenLiteral(), "let")

	letStmt, ok := s.(*ast.LetStatement)
	assert.True(t, ok, fmt.Sprintf("Invalid type assertion with type %T, expected: *ast.LetStatement", s))

	assert.Equal(t, name, letStmt.Name.Value, "Unexpected name value, got %s, expected %s", letStmt.Name.Value, name)
	assert.Equal(t, name, letStmt.Name.TokenLiteral(), "Unexpected token literal, got %s, expected %s", letStmt.TokenLiteral(), name)
}

func TestReturnStatement(t *testing.T) {
	t.Parallel()

	input := `
return 10;
return 5;
return 993322;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	require.NotNil(t, program)

	assert.Len(t, program.Statements, 3, "expecting %d statements, got %d", 3, len(program.Statements))

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		require.True(t, ok)

		assert.Equal(t, "return", returnStmt.TokenLiteral(), fmt.Sprintf("invalid token literal for return statement, got: %s", returnStmt.TokenLiteral()))
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	require.Empty(t, errors, fmt.Sprintf("Errors found in parser: \n%s", strings.Join(errors, "\n")))
}

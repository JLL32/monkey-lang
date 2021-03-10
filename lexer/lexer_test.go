package lexer

import (
	"monkey-lang/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatal("tests[%d] - tokentype wrong. expected=%q, got=%q",
					i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedType {
			t.Fatal("tests[%d] - literal wrong. expected=%q, got=%q",
					i,  tt.expectedLiteral, tok.Literal)
		}
	}

}

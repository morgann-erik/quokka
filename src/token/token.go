package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

/*
 * Token types
 */
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	INDET = "IDENT"
	INT   = "INT"

	ASSIGN = "ASSIGN"
	PLUS   = "PLUS"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

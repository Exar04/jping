package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	COMMA = ","
	COLON = ":"

	LBRACE = "{"
	RBRACE = "}"
	LBRACK = "["
	RBRACK = "]"

	INT    = "INT"
	BOOL   = "BOOL"
	STRING = "STRING"

	TRUE  = "TRUE"
	FALSE = "FALSE"
)

var Keywords = map[string]TokenType{
	"true":  BOOL,
	"false": BOOL,
}

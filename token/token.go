package token

type Type string
type Token struct {
	Type    Type
	Literal string
}

const (
	LET     = "LET"
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	INT     = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	FUNCTION = "FUNCTION"
)

var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(s string) Type {
	if tok, ok := keywords[s]; ok {
		return tok
	}

	return IDENT
}

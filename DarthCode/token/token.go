// token/token.go

package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special tokens
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers & Literals
	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"

	// Operators
	PLUS     TokenType = "+"
	ASSIGN   TokenType = "="
	MINUS    TokenType = "-"
	BANG     TokenType = "!"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"

	LT  TokenType = "<"
	GT  TokenType = ">"
	LTE TokenType = "<="
	GTE TokenType = ">="

	EQ     TokenType = "=="
	NOT_EQ TokenType = "!="

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
)

// keywords is a map for keyword lookups.
var keywords = make(map[string]TokenType, 7)

func init() {
	keywords["fn"] = FUNCTION
	keywords["let"] = LET
	keywords["true"] = TRUE
	keywords["false"] = FALSE
	keywords["if"] = IF
	keywords["else"] = ELSE
	keywords["return"] = RETURN
}

// LookupIdent checks if an identifier is a keyword or a user-defined name.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

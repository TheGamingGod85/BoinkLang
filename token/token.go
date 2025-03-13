// token/token.go

package token

// TokenType represents the type of a token.
type TokenType string

// Token represents a lexical token with a type and literal value.
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
	COLON     TokenType = ":"
	SEMICOLON TokenType = ";"

	LPAREN   TokenType = "("
	RPAREN   TokenType = ")"
	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"
	LBRACKET TokenType = "["
	RBRACKET TokenType = "]"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
	STRING   TokenType = "STRING"

	// Comments
	COMMENT TokenType = "COMMENT"
)

// keywords is a map for keyword lookups.
var keywords = make(map[string]TokenType, 7)

// init initializes the keyword map.
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

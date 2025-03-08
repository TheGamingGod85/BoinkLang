package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special Tokens
	ILLEGAL TokenType = iota
	EOF

	// Identifiers & Literals
	IDENT
	INT

	// Operators
	PLUS
	ASSIGN
	MINUS
	BANG
	ASTERISK
	SLASH

	LT
	GT
	LTE
	GTE

	EQ
	NOT_EQ

	// Delimiters
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
)

// Mapping for quick keyword lookup
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// Lookup table for single-character tokens
var charTokens = map[byte]TokenType{
	'+': PLUS, '-': MINUS, '*': ASTERISK, '/': SLASH,
	'<': LT, '>': GT, '=': ASSIGN, '!': BANG,
	',': COMMA, ';': SEMICOLON, '(': LPAREN, ')': RPAREN,
	'{': LBRACE, '}': RBRACE,
}

// LookupIdent checks if an identifier is a keyword, otherwise returns IDENT
func LookupIdent(ident string) TokenType {
	if tok, found := keywords[ident]; found {
		return tok
	}
	return IDENT
}

package token

type TokenType string

const (
	// Tokens especiales
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Literales
	IDENT  TokenType = "IDENT"
	INT    TokenType = "INT"
	STRING TokenType = "STRING"

	// Operadores
	ASSIGN   TokenType = "="
	CONFIRM  TokenType = ":="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"
	QUESTION TokenType = "?"
	GT       TokenType = ">"
	LT       TokenType = "<"

	// Delimitadores
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"
	LPACK     TokenType = "["
	RPACK     TokenType = "]"

	// Palabras clave
	VAR      TokenType = "VAR"
	FUNC     TokenType = "FUNC"
	STRUCT   TokenType = "STRUCT"
	IMPORT   TokenType = "IMPORT"
	PACK     TokenType = "PACK"
	ECHO     TokenType = "ECHO"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	REQUIRED TokenType = "REQUIRED"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"var":      VAR,
	"func":     FUNC,
	"struct":   STRUCT,
	"import":   IMPORT,
	"pack":     PACK,
	"echo":     ECHO,
	"if":       IF,
	"else":     ELSE,
	"required": REQUIRED,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
} 
package token

// Many diferent values as TokenType, not good for performance but great for debugging
// Using string as TokenType
type TokenType string 

type Token struct {
    Type    TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // Identifiers + literals
    IDENT = "IDENT" // Names
    INT = "INT"     // 123

    // Operators
    ASSIGN = "="
    PLUS = "+"

    // Delimiters
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION    = "FUNCTION"
    LET         = "LET"
)

var keywords = map[string]TokenType{
    "fn": FUNCTION,
    "let": LET,
}

func LookupIdent(ident string) TokenType{
    if tok, ok := keywords[ident]; ok{
        return tok
    }
    return IDENT
}

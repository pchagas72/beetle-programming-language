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
    MINUS = "-"
    BANG = "!"
    ASTERISK = "*"
    SLASH = "/"

    LT = "<"
    GT = ">"

    EQ = "=="
    NOT_EQ = "!="

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
    IF          = "IF"
    ELSE        = "ELSE"
    FALSE       = "FALSE"
    TRUE        = "TRUE"
    RETURN      = "RETURN"
)

var keywords = map[string]TokenType{
    "fn":       FUNCTION,
    "let":      LET,
    "if":       IF,
    "else":     ELSE,
    "false":    FALSE,
    "true":     TRUE,
    "return":   RETURN,
}

func LookupIdent(ident string) TokenType{
    if tok, ok := keywords[ident]; ok{
        return tok
    }
    return IDENT
}

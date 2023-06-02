package token

type TokenType string

type Token struct {
    Type TokenType 
    Literal string
}

var keyword = map[string]TokenType {
    "fn": FUNCTION,
    "let":LET,
    "true":TRUE,
    "false":FALSE,
    "if":IF,
    "else":ELSE,
    "return":RETURN,
    
}
func LookupIdent(ident string) TokenType{
    if tok, ok := keyword[ident]; ok {
        return tok
    }
    return IDENT
}

const (

    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    IDENT   ="IDENT"
    INT     = "INT"

    ASSIGN  = "="
    PLUS    = "+"
    MINUS   = "-"
    BANG    = "!"
    ASTREIK = "*"
    SLASH   = "/"

    LT = "<"
    GT = ">"

    COMMA     = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    FUNCTION = "FUNCTION"
    LET      = "LET"
    TRUE = "TRUE"
    FALSE = "FALSE"
    IF = "IF"
    ELSE = "ELSE"
    RETURN = "RETURN"
)

package parser

import (
    "itpr/token"
    "itpr/lexer"
    "itpr/ast"
) 

type Parser struct {
    l *lexer.Lexer
    
    currToken token.Token
    peekToken token.Token


}
func New(l *lexer.Lexer ) *Parser {
    p := &Parser{l:l}

    //this to initial both curr and peek token 
    p.nextToken()
    p.nextToken()

    return p
}

func (p * Parser) nextToken() {
    p.currToken = p.peekToken
    p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program{
    return nil
}

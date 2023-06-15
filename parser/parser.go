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
    program := &ast.Program{}
    program.Statements = []ast.Statement{}
    
    for p.currToken.Type != token.EOF {
        stmt := p.parseStatement()
        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
        p.nextToken()
    }
    return program
}

func (p *Parser) parseStatement() ast.Statement {
    switch p.currToken.Type {
    case token.LET:
        return p.parseLetStatement()
    default:
        return nil
    }
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
    stmt := &ast.LetStatement{Token: p.currToken}

    //expecting an ident equal sign and last statment until ;

    if !p.expectPeek(token.IDENT){
        return nil
    }
    
    stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
   
    if !p.expectPeek(token.ASSIGN){
        return nil 
    }

    for !p.currTokenIs(token.SEMICOLON){
        p.nextToken()
    }

    return stmt
}
func (p *Parser) currTokenIs(tok token.TokenType) bool {
    return p.currToken.Type == tok
}
func (p *Parser) peekTokenIs(tok token.TokenType) bool {
    return p.peekToken.Type == tok
}

func (p *Parser) expectPeek(tok token.TokenType) bool {
    if p.peekTokenIs(tok){
        p.nextToken()
        return true
    }else {
        return false
    }
    }


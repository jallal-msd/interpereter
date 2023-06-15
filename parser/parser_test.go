package parser

import (
    "testing"
    "itpr/ast"
    "itpr/lexer"
)

func TestLetStatements(t *testing.T){
    input := `
        let x = 5;
        let y = 10;
        let foobar = 8383;
        `
    l := lexer.New(input)
    p :=New(l)

    program :=p.ParseProgram()
    if program == nil {
        t.Fatalf("ParseProgram() returned nil")
    }
    if len(program.Statements) != 3 {
        t.Fatalf("program.statments does not contain 3 statements. got=%d", 
            len(program.Statements))
    } 

    tests :=[]struct {
        expectedIdentifier string
    }{
        {"x"},
        {"y"},
        {"foobar"},
    }

    for i, tt := range tests {
        stmt := program.Statements[i] 
        if !testLetStatement(t,stmt, tt.expectedIdentifier){
            return 
        }
    }
}

func testLetStatement(t *testing.T, s ast.Statement, name string ) bool {
    if s.TokenLiteral() != "let" {
        t.Errorf("s.TokenLiteral not 'let' . got=%q", s.TokenLiteral())
        return false
    }
    
    letStmnt, ok :=s.(*ast.LetStatement) 
    if !ok {
        t.Errorf("s not *ast.LetStatment, got=%T", s)
    }
    if letStmnt.Name.Value != name {
        t.Errorf("letStmnt.Name.Value not '%s'. got=%s",name, letStmnt.Name.Value)
        return false
    }

    if letStmnt.Name.TokenLiteral() != name {
        t.Errorf("s.Name not %s, got=%s", name, letStmnt.Name)
        return false
    }
    return true
}

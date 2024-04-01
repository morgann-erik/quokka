package parser

import (
	"testing"

	"github.com/morgann-erik/quokka/ast"
	"github.com/morgann-erik/quokka/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
    let x=5
    let y = 10
    let foobar = 838383
    `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned null")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements: got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"}, {"y"}, {"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("Expected 'let' token, got %q", s.TokenLiteral())

		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not a *LetStatement, got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("expected TokenLiteral %s, got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("%d parser errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	input := `
    return 7
    return 10
    return 993388
    `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements: got=%d", len(program.Statements))
	}

    for _,stmt := range program.Statements {
        returnStmt, ok := stmt.(*ast.ReturnStatement)
        if !ok {
            t.Errorf("Expected *ast.ReturnStatement, got=%T", stmt)
            continue
        }

        if returnStmt.TokenLiteral() != "return" {
            t.Errorf("Expected 'return', got %q", returnStmt.TokenLiteral())
        }

    }
}

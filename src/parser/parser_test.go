package parser

import (
	"testing"

	"github.com/morgann-erik/quokka/ast"
	"github.com/morgann-erik/quokka/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
    let x=5;
    let y = 10;
    let myVarbar = 838383;
    `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned null")
	}

	if len(program.Statements) != 3 {
        t.Errorf("%s", program.String())
		t.Fatalf("program.Statements does not contain 3 statements: got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"}, {"y"}, {"myVarbar"},
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
    return 7;
    return 10;
    return 993388;
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

func TestString(t *testing.T) {
    input := "let myVar = a;"

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    if program.String() != "let myVar="{
        t.Errorf("program.String() is wrong expected %q, got=%q", "let myVar=", program.String())
    }
}

func TestIdentExpression(t *testing.T) {
    input := "foobar;"

    l := lexer.New(input)
    p := New(l)

    program := p.ParseProgram()
    checkParseErrors(t, p)

    if len(program.Statements) != 1 {
        t.Fatalf("not enough statements expected 1, got %d", len(program.Statements))
    }

    stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
    if !ok {
        t.Fatalf("Expected *ast.ExpressionStatement, got %T", program.Statements[0])
    }

    ident, ok := stmt.Expression.(*ast.Identifier)
    if !ok {
        t.Fatalf("expected Ident expression, got=%T", stmt.Expression)
    }
    if ident.Value != "foobar" {
        t.Errorf("wron ident.Value expected 'foobar', got %s", ident.Value)
    }
    if ident.TokenLiteral() != "foobar" {
        t.Errorf("ident.TokenLiteral not %s, got %s", "foobar", ident.TokenLiteral())
    }
}

func TestIntegerLiteralExpression(t *testing.T) {
    input := "5;"

    l := lexer.New(input)
    p := New(l)
    program := p.ParseProgram()
    checkParseErrors(t, p)

    if len(program.Statements) != 1 {
        t.Fatalf("expected 1 statement, got %d", len(program.Statements))
    }

    stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
    if !ok {
        t.Fatalf("statement is not ExpressionStatement, got %T", program.Statements[0])
    }

    literal, ok := stmt.Expression.(*ast.IntegerLiteral)
    if !ok {
        t.Fatalf("expression is not IntegerLiteral, got %T", stmt.Expression)
    }

    if literal.Value != 5 {
        t.Fatalf("expected literal value of 5, got %d", literal.Value)
    }

    if literal.TokenLiteral() != "5" {
        t.Fatalf("expected TokenLiteral of 5, got %s", literal.TokenLiteral())
    }
}

package parser

import (
	"interpreterInGo/ast"
	"interpreterInGo/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 4859834;
`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatalf("falal error: returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contailn 3 statements. Got :%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
return 100;
return 1;
return 1003234;
`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatalf("falal error: returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contailn 3 statements. Got :%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnstmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			//
		}
		if returnstmt.TokenLiteral() != "return" {
			t.Errorf("returnstmt.TokenLiteral() not 'return', got %q", returnstmt.TokenLiteral())
		}
	}
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.errors
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.Tokenliteral is not 'let'. got: %q", s.TokenLiteral())
		return false
	}
	letstmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s is not '*ast.LetStatement'. got: %T", s)
		return false
	}

	if letstmt.Name.Value != name {
		t.Errorf("letstmt.Name.Value is not %s. got: %s", name, letstmt.Name.Value)
		return false
	}

	if letstmt.Name.TokenLiteral() != name {
		t.Errorf("letstmt.Name.TokenLiteral() is not %s. got: %s", name, letstmt.Name.TokenLiteral())
		return false
	}

	return true
}

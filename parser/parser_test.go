package parser

import (
	"testing"

	"github.com/carepollo/sexlang/ast"
	"github.com/carepollo/sexlang/lexer"
)

func TestLetStatement(t *testing.T) {
	input := `foobar;`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("Parse program() returned nil")
	}
	if len(program.Statements) != 1 {
		// fmt.Println(*(program).Statements)
		t.Fatalf("program.Statements does not contain 3 statements. got %d", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not as.ExpressionStatement, got %T", program.Statements[0])
	}

	ident, ok := statement.Expression.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier, got %T", statement.Expression)
	}
	if ident.Value != "foobar" {
		t.Errorf("ident.Value %s, got %s", "foobar", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s not %s", "foobar", ident.TokenLiteral())
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errores", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

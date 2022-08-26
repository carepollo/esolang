package evaluator

import (
	"github.com/carepollo/sexlang/ast"
	"github.com/carepollo/sexlang/object"
)

// decide which evaluate which value according the the type node
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.Boolean:
		return &object.Boolean{
			Value: node.Value,
		}
	case *ast.IntegerLiteral:
		return &object.Integer{
			Value: node.Value,
		}
	}
	return nil
}

// recursive way to keep evaluation of statements
func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range statements {
		result = Eval(statement)
	}
	return result
}

package evaluator

import (
	"fmt"

	"github.com/carepollo/sexlang/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Function: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError(
					"wrong number of arguments. got=%d, want=1",
					len(args),
				)
			}

			switch argument := args[0].(type) {
			case *object.Array:
				return &object.Integer{
					Value: int64(len(argument.Elements)),
				}
			case *object.String:
				return &object.Integer{
					Value: int64(len(argument.Value)),
				}
			default:
				return newError(
					"argument to `len` not supported, got %s",
					args[0].Type(),
				)
			}
		},
	},
	"first": {
		Function: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments, got %d want 1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'first' must be ARRAY, got %s", args[0].Type())
			}

			array := args[0].(*object.Array)
			if len(array.Elements) > 0 {
				return array.Elements[0]
			}

			return NULL
		},
	},
	"last": {
		Function: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments, got %d want 1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'first' must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"rest": {
		Function: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments, got %d want 1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to 'first' must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{
					Elements: newElements,
				}
			}
			return NULL
		},
	},
	"push": {
		Function: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			newElements := make([]object.Object, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]
			return &object.Array{Elements: newElements}
		},
	},
	"puts": {
		Function: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
}

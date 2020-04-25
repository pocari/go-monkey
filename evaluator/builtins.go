package evaluator

import "monkey/object"

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: builtinLen,
	},
	"first": &object.Builtin{
		Fn: builtinFirst,
	},
	"last": &object.Builtin{
		Fn: builtinLast,
	},
	"rest": &object.Builtin{
		Fn: builtinRest,
	},
}

func builtinLen(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{
			Value: int64(len(arg.Value)),
		}
	case *object.Array:
		return &object.Integer{
			Value: int64(len(arg.Elements)),
		}
	default:
		return newError("argument to `len` not supported, got %s", args[0].Type())
	}
}

func builtinFirst(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.Array:
		return arg.Elements[0]
	default:
		return newError("argument to `first` not supported, got %s", args[0].Type())
	}
}

func builtinLast(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.Array:
		return arg.Elements[len(arg.Elements)-1]
	default:
		return newError("argument to `last` not supported, got %s", args[0].Type())
	}
}

func builtinRest(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.Array:
		length := len(arg.Elements)
		if length > 0 {
			newElements := make([]object.Object, length-1, length-1)
			copy(newElements, arg.Elements[1:length])
			return &object.Array{
				Elements: newElements,
			}
		}
		return NULL
	default:
		return newError("argument to `rest` not supported, got %s", args[0].Type())
	}
}

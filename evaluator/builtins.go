package evaluator

import (
	"fmt"
	"monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len":   {Fn: builtinLen},
	"first": {Fn: builtinFirst},
	"last":  {Fn: builtinLast},
	"push":  {Fn: builtinPush},
	"rest":  {Fn: builtinRest},
	"puts":  {Fn: builtinPuts},
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

func builtinPush(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2", len(args))
	}

	switch arg := args[0].(type) {
	case *object.Array:
		length := len(arg.Elements)
		newElements := make([]object.Object, length+1, length+1)
		copy(newElements, arg.Elements)
		newElements[length] = args[1]
		return &object.Array{
			Elements: newElements,
		}
	default:
		return newError("argument to `push` not supported, got %s", args[0].Type())
	}
}

func builtinPuts(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}
	return NULL
}

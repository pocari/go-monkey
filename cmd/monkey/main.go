package main

import (
	"fmt"
	"io/ioutil"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/repl"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("input file required")
	}

	inputFile := os.Args[1]
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	l := lexer.New(string(bytes))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		repl.PrintParserError(os.Stdout, p.Errors())
		return
	}

	env := object.NewEnvironment()
	evaluator.Eval(program, env)
}

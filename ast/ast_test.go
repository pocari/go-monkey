package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	// let myVar = anotherVar;
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: &token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: &token.Token{Type: token.IDENT, Literal: "IDENT"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: &token.Token{Type: token.IDENT, Literal: "IDENT"},
					Value: "anotherValue",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherValue;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

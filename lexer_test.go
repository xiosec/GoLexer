package golexer

import "testing"

func TestLexer(t *testing.T) {
	ruls := [][]string{
		{`\d+`, "NUMBER"},
		{`[a-zA-Z_]\w+`, "IDENTIFIER"},
		{"=", "EQUALS"},
	}

	variable := Lexer{
		Rules: ruls,
	}
	variable.Init()
	if variable.Input("age=12")[0].Type != "IDENTIFIER" {
		t.Error("age is Identifier")
	}
}

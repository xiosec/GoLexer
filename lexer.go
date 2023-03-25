package golexer

import (
	"fmt"
	"regexp"
	"strings"
)

type Token struct {
	Type  string
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf("%s\t%s", t.Value, t.Type)
}

type Lexer struct {
	Rules [][]string

	group map[string]string
	re    regexp.Regexp
}

func (lex *Lexer) Init() {
	idx := 1
	var parts []string
	lex.group = make(map[string]string)

	for _, rul := range lex.Rules {
		regex := rul[0]
		name := rul[1]

		groupname := fmt.Sprintf("GROUP%d", idx)
		parts = append(parts, fmt.Sprintf("(?P<%s>%s)", groupname, regex))
		lex.group[groupname] = name
		idx += 1
	}
	re := regexp.MustCompile(strings.Join(parts, "|"))
	lex.re = *re
}

func (lex *Lexer) Input(buff string) []Token {
	groupNames := lex.re.SubexpNames()
	var tokens []Token

	for _, match := range lex.re.FindAllStringSubmatch(buff, -1) {
		for Idx, value := range match {
			name := lex.group[groupNames[Idx]]
			if value != "" && name != "" {
				tokens = append(tokens, Token{
					Type:  name,
					Value: value,
				})
			}
		}
	}
	return tokens
}

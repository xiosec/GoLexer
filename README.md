# GoLexer
Example of Lexer for golang (based on RegEx)

# Example
```go
//A list of rules. Each rule is a `regex, type`
//pair, where `regex` is the regular expression used
//to recognize the token and `type` is the type
//of the token to return when it's recognized.

ruls := [][]string{
	{`\d+`, "NUMBER"},
	{`[a-zA-Z_]\w+`, "IDENTIFIER"},
	{`\+`, "PLUS"},
	{`\-`, "MINUS"},
	{`\*`, "MULTIPLY"},
	{`\/`, "DIVIDE"},
	{`\(`, "LP"},
	{`\)`, "RP"},
	{"=", "EQUALS"},
}

variable := Lexer{
		Rules: ruls,
}

variable.Init()
str := "age = (now - birth) + 1 / (2 * 2)"

for _, token :=range variable.Input(str){
    fmt.Println(token)
}
```
**output**
```shell
Value   Type

age     IDENTIFIER
=       EQUALS
(       LP
now     IDENTIFIER
-       MINUS
birth   IDENTIFIER
)       RP
+       PLUS
1       NUMBER
/       DIVIDE
(       LP
2       NUMBER
*       MULTIPLY
2       NUMBER
)       RP
```
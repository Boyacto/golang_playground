package main

import (
	"fmt"

	gen "patito/gen"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	source := `
program control ;
vars x : int ;
main {
x = 10 ;
if ( x > 5 ) { print(x) ; } else { x = 0 ; } ;
while ( x > 0 ) do { x = x - 1 ; } ;
}
end
`
	input := antlr.NewInputStream(source)
	lex := gen.NewPatitoLexer(input)
	tokens := antlr.NewCommonTokenStream(lex, 0)
	p := gen.NewPatitoParser(tokens)

	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := p.Program()
	fmt.Println(tree.ToStringTree(nil, p))
}

package main

import (
	"fmt"

	gen "patito/gen"
	"patito/semantic"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	// Example Patito program with semantic analysis
	source := `
program ejemplo ;
vars
	x, y : int ;
	resultado : float ;

void sumar(a : int, b : int) {
	vars temp : int ;
	{
		temp = a + b ;
		resultado = temp ;
	}
}
;

main {
	x = 10 ;
	y = 20 ;
	sumar(x, y) ;
	print(resultado) ;
}
end
`

	fmt.Println("=== Patito Semantic Analyzer ===")
	fmt.Println("\nSource Code:")
	fmt.Println(source)

	// Create lexer and parser
	input := antlr.NewInputStream(source)
	lexer := gen.NewPatitoLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	parser := gen.NewPatitoParser(tokens)

	// Remove default error listeners
	parser.RemoveErrorListeners()

	// Parse the program
	tree := parser.Program()

	fmt.Println("\n=== Syntax Analysis ===")
	fmt.Println("Parse tree created successfully")

	// Create semantic analyzer
	analyzer := semantic.NewSemanticAnalyzer()

	// Walk the parse tree with the semantic analyzer
	antlr.ParseTreeWalkerDefault.Walk(analyzer, tree)

	// Print results
	fmt.Println("\n=== Semantic Analysis ===")

	if analyzer.HasErrors() {
		fmt.Println("❌ Semantic errors found:")
		for _, err := range analyzer.GetErrors() {
			fmt.Printf("  - %s\n", err)
		}
	} else {
		fmt.Println("✓ No semantic errors found!")
	}

	// Print symbol tables
	fmt.Println("\n=== Global Variables ===")
	vm := analyzer.GetVariableManager()
	globalTable := vm.GetGlobalTable()
	for name, symbol := range globalTable.GetAll() {
		fmt.Printf("  %s: %s (address: %d)\n", name, symbol.Type, symbol.Address)
	}

	// Print function directory
	fmt.Println("\n=== Functions ===")
	fd := analyzer.GetFunctionDirectory()
	for name, funcSig := range fd.GetAll() {
		fmt.Printf("  %s(", name)
		for i, param := range funcSig.Parameters {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%s: %s", param.Name, param.Type)
		}
		fmt.Printf(") -> %s\n", funcSig.ReturnType)

		// Print local variables
		if funcSig.LocalVarTable.Size() > 0 {
			fmt.Printf("    Local variables:\n")
			for localName, localSymbol := range funcSig.LocalVarTable.GetAll() {
				fmt.Printf("      %s: %s (address: %d)\n", localName, localSymbol.Type, localSymbol.Address)
			}
		}
	}
}

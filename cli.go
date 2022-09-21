package main

import (
	"fmt"

	mathematics "github.com/masc23/go-algorithms-echo/mathematics/equations"
)

func main() {
	//testEquations := [3]string{
	//	"1.2 + 2,3 - 3 * 4 / 5 * (6 + 7 / (8 + 9)",
	//	"f(x) = x^2 + 2x - 3",
	//	"cos(x)^2 * sin(x) + 2x",
	//}

	testEquations := [1]string{
		"1 + 2 * 3",
	}

	for _, s := range testEquations {
		r := mathematics.Lexer(s)
		p, err := mathematics.Parser(r)

		if err != nil {
			fmt.Println("ERROR: " + err.Error())
			continue
		}

		//testLexer(s, r)
		testParser(s, r, p)
		//testEvaluator(s, r, p)
	}
}

//lint:ignore U1000 only used for development
func testLexer(s string, r []mathematics.Token) {
	fmt.Printf("Original Equation      = %s\n", s)
	fmt.Printf("Reconstructed Equation = %s\n", mathematics.PrintEquation(r))
	fmt.Printf("Number of Tokens       = %d\n", len(r))
	fmt.Println()

	fmt.Println(mathematics.PrettyPrintTokenList(r))
	fmt.Println()
}

//lint:ignore U1000 only used for development
func testParser(s string, r []mathematics.Token, p mathematics.Equation) {
	fmt.Printf("parsing %s\n", mathematics.PrintEquation(r))
	fmt.Println()
}

/*
//lint:ignore U1000 only used for development
func testEvaluator(s string, r []mathematics.Token, p mathematics.Evaluable) {
	fmt.Printf("evaluating %s\n", mathematics.PrintEquation(r))
	fmt.Printf("  result = %.5f", p.Evaluate())
	fmt.Println()
	fmt.Println()
}
*/

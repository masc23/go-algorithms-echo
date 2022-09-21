package mathematics

import (
	"fmt"
	"strings"
)

func PrintEquation(tokens []Token) string {
	sb := new(strings.Builder)

	for _, t := range tokens {
		sb.WriteString(t.value)
	}

	return sb.String()
}

func PrettyPrintTokenList(tokens []Token) string {
	sb := new(strings.Builder)

	sb.WriteString(fmt.Sprintf("%-20s %s\n", "Type", "Value"))
	sb.WriteString(fmt.Sprintln("--------------------------"))

	for _, t := range tokens {
		sb.WriteString(fmt.Sprintf("%-20s %s\n", t.tokenType, t.value))
	}

	sb.WriteRune('\n')

	return sb.String()
}

func PrettyPrintSyntaxTree() string {
	sb := new(strings.Builder)

	return sb.String()
}

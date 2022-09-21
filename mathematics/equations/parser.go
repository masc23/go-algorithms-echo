package mathematics

type Equation interface{}
type Term interface{}

type ParserError struct {
	Msg string
}

func (pe ParserError) Error() string {
	return pe.Msg
}

type FullEquation struct {
	functionName string
	variableList []string
	equationTerm Term
	Equation
}

type SimpleEquation struct {
	equationTerm Term
	Equation
}

type Number struct {
	value float64
	Term
}

type identifier struct {
	value string
	Term
}

func Parser(tokenList []Token) (Equation, error) {
	var detailedErrorMsg string

	// at first, try to parse a full equation
	if eq, err := parseFullEquation(tokenList); err == nil {
		return eq, nil
	}

	// then try to parse a simple equation
	if eq, err := parseSimpleEquation(tokenList); err == nil {
		return eq, nil

	} else {
		detailedErrorMsg = err.Error()
	}

	return nil, ParserError{Msg: "tokenList does not resemble neither a full nor a simple equation; " + detailedErrorMsg}
}

func parseFullEquation(tokenList []Token) (*FullEquation, error) {
	return nil, ParserError{Msg: "parseFullEquation not implemented yet"}
}

func parseSimpleEquation(tokenList []Token) (*SimpleEquation, error) {
	// a simple equation is only a term, so try to parse that
	eq, err := parseTerm(tokenList)

	if err != nil {
		return nil, err
	}

	return &SimpleEquation{equationTerm: eq}, nil
}

func parseTerm(tokenList []Token) (Term, error) {
	return parseNumber(tokenList)
}

func parseNumber(tokenList []Token) (*Number, error) {
	n := Number{}

	return &n, nil
}

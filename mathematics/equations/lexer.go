package mathematics

type Token struct {
	tokenType string
	value     string
}

// Takes an input string and converts it to a list of
// tokens. Does no checking for syntax or semantics.
func Lexer(input string) []Token {
	runes := make(chan rune)
	tokens := make(chan Token)

	go disassembleString(runes, input)
	go tokenize(runes, tokens)

	tokenList := make([]Token, 0, 100)

	for token := range tokens {
		tokenList = append(tokenList, token)
	}

	eofToken := Token{tokenType: "eof", value: ""}

	return append(tokenList, eofToken)
}

// Disassembles the input string into single runes and writes
// them to the channel runes.
// Closes the channels after the string is completely read.
func disassembleString(runes chan rune, input string) {
	defer close(runes)

	for _, c := range input {
		runes <- c
	}
}

// Reads runes from the channel runes and tries to convert them
// to actual tokens. Each token is then written to the channel
// tokens.
// Closes the channel tokens after the channel runes is closed.
func tokenize(runes chan rune, tokens chan Token) {
	defer close(tokens)

	lexing := true
	carryOver := false

	var c rune
	var token Token
	var notEndOfFile bool

	for lexing {
		if !carryOver {
			c, notEndOfFile = <-runes

			if !notEndOfFile {
				lexing = false
			}
		} else {
			carryOver = false
		}

		if isDigit(c) {
			// found a digit, continue to parse the whole number
			token, c = createNumberToken(c, runes)
			tokens <- token
			carryOver = true

		} else if isAlpha(c) {
			// found a letter, continue to parse the whole string
			// note: in equations like 'xy - 2', x and y are two
			// different identifiers
			// this situation is handled by the parser
			token, c = createStringToken(c, runes)
			tokens <- token
			carryOver = true

		} else if isOperator(c) {
			// found an operator
			tokens <- Token{tokenType: "operator", value: string(c)}

		} else if c == '(' {
			// found an opening parenthesis
			tokens <- Token{tokenType: "openingParenthesis", value: string(c)}

		} else if c == ')' {
			// found an closing parenthesis
			tokens <- Token{tokenType: "closingParenthesis", value: string(c)}

		} else if c == '=' {
			// found an equals sign
			tokens <- Token{tokenType: "equals", value: string(c)}

		} else if isWhitespace(c) {
			// just skip whitespaces

		} else {
			// found a non-valid token
			tokens <- Token{tokenType: "invalid", value: string(c)}
		}
	}
}

// Reads from the channel as long as there are digits, ',' or '.'
// and returns a constructed number token.
// Since the first rune is already read from the channel, it is
// also supplied and used as the first rune inside the token value.
// The last read rune, which is not a part of the token, is
// also returned.
func createNumberToken(firstRune rune, runes chan rune) (token Token, nextRune rune) {
	token = Token{tokenType: "number", value: string(firstRune)}

	lexing := true

	for lexing {
		nextRune = <-runes

		if isDigit(nextRune) || nextRune == '.' || nextRune == ',' {
			token.value += string(nextRune)
		} else {
			lexing = false
		}
	}

	return
}

// Reads from the channel as long as there are alphas
// and returns a constructed string token.
// Since the first rune is already read from the channel, it is
// also supplied and used as the first rune inside the token value.
// The last read rune, which is not a part of the token, is
// also returned.
func createStringToken(firstRune rune, runes chan rune) (token Token, nextRune rune) {
	token = Token{tokenType: "string", value: string(firstRune)}

	lexing := true

	for lexing {
		nextRune = <-runes

		if isAlpha(nextRune) {
			token.value += string(nextRune)
		} else {
			lexing = false
		}
	}

	return
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isAlpha(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

func isOperator(c rune) bool {
	return c == '+' || c == '-' || c == '*' || c == '/' || c == '^'
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == 0
}

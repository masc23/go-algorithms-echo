package rot13

import "strings"

func rot13String(input string) string {
	sb := new(strings.Builder)

	for _, c := range input {
		character := int(c)

		// move 13 steps
		rotation := 13

		if character >= 'A' && character <= 'Z' {
			// wrap-around
			if character+rotation > 'Z' {
				rotation -= 26
			}

		} else if character >= 'a' && character <= 'z' {
			// wrap-around
			if character+rotation > 'z' {
				rotation -= 26
			}

		} else {
			// reset the rotation for non A-Za-z characters
			rotation = 0
		}

		character += rotation

		sb.WriteRune(rune(character))
	}

	return sb.String()
}

package base64

import (
	"strconv"
	"strings"
)

func decodeString(input string) string {
	sb := new(strings.Builder)

	for _, c := range input {
		sb.WriteString(FromBase64(string(c)))
	}

	binaryString := sb.String()
	sb.Reset()

	originalStringLength := len(binaryString) - strings.Count(input, "=")*8

	for i := 0; i < originalStringLength; i += 8 {
		sixBitString := binaryString[i : i+8]
		block, _ := strconv.ParseInt(sixBitString, 2, 0)
		sb.WriteRune(rune(block))
	}

	return sb.String()
}

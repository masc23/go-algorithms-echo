package base64

import (
	"strconv"
	"strings"
)

func encodeString(input string) string {
	binaryString := convertStringToBinaryString(input)

	fillZeroes := 0
	for (len(binaryString)/8)%3 != 0 {
		binaryString += "0"
		fillZeroes++
	}
	padding := fillZeroes / 6

	sb := new(strings.Builder)

	for i := 0; i < len(binaryString)-fillZeroes; i += 6 {
		sixBitString := binaryString[i : i+6]
		block, _ := strconv.ParseInt(sixBitString, 2, 0)
		sb.WriteString(ToBase64(int(block)))
	}

	for i := 0; i < padding; i++ {
		sb.WriteString("=")
	}

	return sb.String()
}

func convertStringToBinaryString(input string) string {
	sb := new(strings.Builder)

	for _, c := range input {
		sb.WriteString(ToPaddedBinary(int(c), 8))
	}

	return sb.String()
}

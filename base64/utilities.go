package base64

import (
	"fmt"
)

var b64 = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "+", "/"}

func ToBase64(idx int) string {
	return b64[idx]
}

func FromBase64(s string) string {
	idx := 0
	for i, b := range b64 {
		if b == s {
			idx = i
			break
		}
	}

	return ToPaddedBinary(idx, 6)
}

func ToPaddedBinary(n int, padding int) string {
	binary := fmt.Sprintf("%b", n)

	for len(binary) < padding {
		binary = "0" + binary
	}

	return binary
}

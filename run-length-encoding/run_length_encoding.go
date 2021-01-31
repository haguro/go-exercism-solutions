//Package encode provides the functionality to encode and decode strings using
//using the Run Length algorithms.
package encode

import (
	"strconv"
	"unicode"
)

//RunLengthEncode encodes a sequence of characters using RLE.
func RunLengthEncode(input string) string {
	var encoded string
	runes := []rune(input)
	count := 1
	length := len(runes)
	for i, r := range runes {
		if i < length-1 && r == runes[i+1] {
			count++
		} else {
			if count > 1 {
				encoded += strconv.Itoa(count)
			}
			encoded += string(r)
			count = 1
		}
	}
	return encoded
}

//RunLengthDecode decodes a string encoded with RLE to the original sequence
//of characters.
func RunLengthDecode(input string) string {
	var digits []rune
	var decoded string
	count := 1
	for _, r := range input {
		if unicode.IsDigit(r) {
			digits = append(digits, r)
			continue
		}
		if len(digits) > 0 {
			count, _ = strconv.Atoi(string(digits))
		}
		for i := 0; i < count; i++ {
			decoded += string(r)
		}
		digits = []rune{}
		count = 1
	}
	return decoded
}

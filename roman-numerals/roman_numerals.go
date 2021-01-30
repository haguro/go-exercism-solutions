// Package romannumerals provides functionality to convert Arabic numbers to
// Roman numerals.
package romannumerals

import (
	"errors"
)

var symbolValues = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var symbolMap = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

// ToRomanNumeral convert an integer between 1 and 3000 to a roman numeral equivalent.
func ToRomanNumeral(arabic int) (string, error) {
	if arabic <= 0 || arabic > 3000 {
		return "", errors.New("arabic number out of bound")
	}
	var roman string
	for _, v := range symbolValues {
		for arabic >= v {
			roman += symbolMap[v]
			arabic -= v
		}
	}
	return roman, nil
}

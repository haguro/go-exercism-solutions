package lsproduct

import (
	"errors"
)

//LargestSeriesProduct returns the the largest product of multiplying a given span of consecutive digits.
func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) || span < 0 {
		return 0, errors.New("span is larger than length of string or less than 0")
	}
	var largest int64
	for i := 0; i <= len(digits)-span; i++ {
		var product int64 = 1
		for j := 0; j < span; j++ {
			digit := digits[i+j]
			if digit < 0x30 || digit > 0x39 {
				return 0, errors.New("string contains non-digit characters")
			}
			product *= int64(digit - 0x30)
		}
		if largest < product {
			largest = product
		}
	}
	return largest, nil
}

// Package leap provides functionality for leap year detection.
package leap

// IsLeapYear returns true if a given year isa leap year
func IsLeapYear(year int) bool {
	if (year%4 == 0) && ((year%100 != 0) || (year%400 == 0)) {
		return true
	}
	return false
}

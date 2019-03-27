//Package etl provides the functionality to convert legacy scrabble score maps to the new and shiny (also more logical) format.
package etl

import "strings"

//Transform coverts legacy score map format to the new format
func Transform(oldMap map[int][]string) map[string]int {
	newMap := make(map[string]int)
	for score, letters := range oldMap {
		for _, letter := range letters {
			newMap[strings.ToLower(letter)] = score
		}
	}
	return newMap
}

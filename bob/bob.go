//Package bob provides functionality to simulate a lackadaisical teenager's responses.
package bob

import (
	"strings"
)

func isAllCaps(s string) bool {
	return strings.ToUpper(s) == s && strings.ContainsAny(s, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

//Hey returns the teenager's responses based on a given communcation.
func Hey(remark string) string {
	s := strings.TrimSpace(remark)
	switch {
	case len(s) == 0:
		return "Fine. Be that way!"
	case isAllCaps(s) && strings.HasSuffix(s, "?"):
		return "Calm down, I know what I'm doing!"
	case strings.HasSuffix(s, "?"):
		return "Sure."
	case isAllCaps(s):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

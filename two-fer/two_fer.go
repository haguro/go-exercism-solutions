// Package twofer provides the ultimate answer of what to do when you have two of something.
package twofer

import "fmt"

// ShareWith returns the two-fer text optionally based on a given name
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

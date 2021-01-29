// Package clock provides functionality to represent, add to and subtract from clock time.
package clock

import (
	"fmt"
)

// Clock represent clock time.
type Clock struct {
	m int
}

// New returns a new value of type Clock ensuring proper roll over
// for out of bound hour and minute values.
func New(h, m int) Clock {
	m += h * 60
	m %= 24 * 60
	if m < 0 {
		m += 24 * 60
	}
	return Clock{m}
}

// String returns string representation of a Clock value.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.m/60, c.m%60)
}

// Add adds minutes to a Clock value and returns a Clock with
// the new time.
func (c Clock) Add(m int) Clock {
	return New(c.m/60, m+c.m%60)
}

// Subtract subtract minutes from a Clock value and returns a Clock
// with the new time.
func (c Clock) Subtract(m int) Clock {
	return c.Add(-m)
}

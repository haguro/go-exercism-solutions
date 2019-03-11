// Package gigasecond provide functionality to calculate the moment when someone turns 10^9 seconds old
package gigasecond

import "time"

// AddGigasecond adds 10^9 seconds to a given date and returns the result
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}

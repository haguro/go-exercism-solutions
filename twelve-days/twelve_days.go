package twelve

import (
	"fmt"
	"strings"
)

const template = "On the %s day of Christmas my true love gave to me: %s."

var numbers = []struct {
	cardinal, ordinal string
}{
	{"a", "first"},
	{"two", "second"},
	{"three", "third"},
	{"four", "fourth"},
	{"five", "fifth"},
	{"six", "sixth"},
	{"seven", "seventh"},
	{"eight", "eighth"},
	{"nine", "ninth"},
	{"ten", "tenth"},
	{"eleven", "eleventh"},
	{"twelve", "twelfth"},
}

var gifts = []string{
	"Partridge in a Pear Tree",
	"Turtle Doves",
	"French Hens",
	"Calling Birds",
	"Gold Rings",
	"Geese-a-Laying",
	"Swans-a-Swimming",
	"Maids-a-Milking",
	"Ladies Dancing",
	"Lords-a-Leaping",
	"Pipers Piping",
	"Drummers Drumming",
}

//Song combines all twelves verses into a single string
func Song() string {
	var vs []string
	for i := 1; i <= 12; i++ {
		vs = append(vs, Verse(i))
	}
	return strings.Join(vs, "\n")
}

//Verse generates a verse that accumulates gifts based on the day of the day of the day
func Verse(n int) string {
	if n < 1 {
		return ""
	}
	var b strings.Builder

	for i := n; i > 0; i-- {
		b.WriteString(numbers[i-1].cardinal + " " + gifts[i-1])
		if i > 1 {
			b.WriteString(", ")
		}
		if n > 1 && i == 2 {
			b.WriteString("and ")
		}
	}

	return fmt.Sprintf(template, numbers[n-1].ordinal, b.String())
}

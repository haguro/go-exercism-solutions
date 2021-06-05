package twelve

import (
	"fmt"
	"strings"
)

const template = "On the %s day of Christmas my true love gave to me: %s."

var ordinals = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var gifts = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
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

	for i := 1; i < n; i++ {
		b.WriteString(gifts[n-i])
		b.WriteString(", ")
	}

	if n > 1 {
		b.WriteString("and ")
	}
	b.WriteString(gifts[0])

	return fmt.Sprintf(template, ordinals[n-1], b.String())
}

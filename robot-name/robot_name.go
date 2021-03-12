package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot represents a robot.
type Robot struct {
	name string
}

const maxUsedNames = 26 * 26 * 10 * 10 * 10

var usedNames = map[string]bool{}

// Name returns the robot name if it exists or generates a new one if it doesn't.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(usedNames) > maxUsedNames {
		return "", errors.New("robot namespace exausted")
	}
	name := generateName()
	for usedNames[name] {
		name = generateName()
	}
	r.name = name
	usedNames[r.name] = true

	return r.name, nil
}

// Reset resets the robot's name
func (r *Robot) Reset() {
	r.name = ""
}

func generateName() string {
	rand.Seed(time.Now().UnixNano())
	char := func() rune {
		return 'A' + rune(rand.Intn(26))
	}
	return fmt.Sprintf("%c%c%03d", char(), char(), rand.Intn(1000))
}

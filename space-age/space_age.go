//Package space provides the ability to calculate someone age on a given
//planet (of Sol).
package space

//Planet represent a Solar system planet
type Planet string

const earthOp = 31557600

// Orbital periods in seconds relative to Earth's
var opr = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

//Age returns the age in seconds on a given Solar system planet.
func Age(seconds float64, planet Planet) float64 {
	return seconds / (earthOp * opr[planet])
}

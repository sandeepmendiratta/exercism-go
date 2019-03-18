package space

import "fmt"

var secondsInOneEarthYear = 31557600.0
var orbitalPeriodsInEarthYears = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

type Planet string

//Age function to convert sec to year
func Age(seconds float64, planet Planet) float64 {

	secondsInAYearOnPlanet := secondsInOneEarthYear * orbitalPeriodsInEarthYears[planet]

	years := seconds / secondsInAYearOnPlanet

	return years

}
func main() {
	fmt.Println("Earth:", Age(1000000, "Earth"))
}

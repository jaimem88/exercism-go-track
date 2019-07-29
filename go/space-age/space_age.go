package space

// Planet string alias
type Planet string

// nolint
const (
	oneEarthYearSeconds float64 = 31557600
	Mercury             Planet  = "Mercury"
	Venus               Planet  = "Venus"
	Earth               Planet  = "Earth"
	Mars                Planet  = "Mars"
	Jupiter             Planet  = "Jupiter"
	Saturn              Planet  = "Saturn"
	Uranus              Planet  = "Uranus"
	Neptune             Planet  = "Neptune"
)

var earthYearsPerPlanet = map[Planet]float64{
	Mercury: 0.2408467,
	Venus:   0.61519726,
	Earth:   1,
	Mars:    1.8808158,
	Jupiter: 11.862615,
	Saturn:  29.447498,
	Uranus:  84.016846,
	Neptune: 164.79132,
}

// Age calculates the age in earth years based on seconds per planet
func Age(seconds float64, planet Planet) float64 {
	planetEarthYears, ok := earthYearsPerPlanet[planet]
	if !ok {
		return 0
	}

	return seconds / (planetEarthYears * oneEarthYearSeconds)
}

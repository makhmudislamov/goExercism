package space

// formula = Age in seconds / (planet orbital period in Earth years * 31557600 (earth year in seconds))

// import "math"


// Age calculates age in seconds on different planets
func Age(seconds float64, planet string) float64 {
	
	// PSEUDOCODE
	// Input >> planet string and second float64
	// Output >> float64
	// 
	
	
	if planet == "Earth" {

		return seconds / 31557600
	} else {

		return 123
	}

	
}
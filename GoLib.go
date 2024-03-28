package GoLib

import (
	"math"
)


func Round(f float64, d float64) float64 {
	var h float64 = math.Pow(10,d)
	return math.Round(f*h)/h
}

// have to add "to decimal places" functionality
func RoundUp(f) {
	return math.Ceil(f)	
}

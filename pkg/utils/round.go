package utils

import "math"

var (
	kilo = math.Pow(10, 3)
	mega = math.Pow(10, 6)
	giga = math.Pow(10, 9)
	tera = math.Pow(10, 12)
	peta = math.Pow(10, 15)
)

func roundOffNearestTen(num float64, divisor float64) float64 {
	x := num / divisor
	return math.Round(x*10) / 10
}

// RoundValues rounds off a pair of given floats to Thousands (kilo),
// Millions (mega) or Billions (giga).
func RoundValues(num1, num2 float64) ([]float64, string) {
	nums := []float64{}
	var units string
	var n float64
	if num1 > num2 {
		n = num1
	} else {
		n = num2
	}

	switch {
	case n < kilo:
		nums = append(nums, num1)
		nums = append(nums, num2)
		units = ""

	case n < mega:
		nums = append(nums, roundOffNearestTen(num1, kilo))
		nums = append(nums, roundOffNearestTen(num2, kilo))
		units = "K"

	case n < giga:
		nums = append(nums, roundOffNearestTen(num1, mega))
		nums = append(nums, roundOffNearestTen(num2, mega))
		units = "M"

	case n < tera:
		nums = append(nums, roundOffNearestTen(num1, giga))
		nums = append(nums, roundOffNearestTen(num2, giga))
		units = "B"

	case n >= peta:
		nums = append(nums, roundOffNearestTen(num1, tera))
		nums = append(nums, roundOffNearestTen(num2, tera))
		units = "T"

	default:
		return []float64{num1, num2}, ""
	}

	return nums, units
}

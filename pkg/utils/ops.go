package utils

// MinFloat64 returns minimum float from a given number of floats
func MinFloat64(a ...float64) float64 {
	var min float64
	if len(a) > 0 {
		min = a[0]
	} else {
		return 0
	}

	for _, val := range a {
		if val < min {
			min = val
		}
	}
	return min
}

// MaxFloat64 returns maximum float from a given number of floats
func MaxFloat64(a ...float64) float64 {
	var max float64
	if len(a) > 0 {
		max = a[0]
	} else {
		return 0
	}
	for _, val := range a {
		if val > max {
			max = val
		}
	}
	return max
}

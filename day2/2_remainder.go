package main

import "math"

func quotientReverse(x int) int {

	y := 0
	r := 0
	sign := 1

	if x < 0 {
		sign = -1
		x = -x
	}

	for x > 0 {
		r = x % 10
		y = y*10 + r
		x = x / 10
	}

	y *= sign

	if (y > math.MaxInt32) || (y < math.MinInt32) {
		return 0
	}

	return y
}

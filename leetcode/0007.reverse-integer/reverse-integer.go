package problem0007

import "math"

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = x * sign
	}
	var ret int
	for x > 0 {
		ret = ret*10 + x%10
		x = x / 10
	}

	ret = ret * sign

	if ret > math.MaxInt32-1 || ret < math.MinInt32 {
		ret = 0
	}

	return ret
}

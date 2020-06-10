package problem0011

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		area := (right - left) * min(height[left], height[right])
		if area > max {
			max = area
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}

	}
	return max
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

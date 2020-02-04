package problem0016

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minSub := math.MaxInt64
	var retArry [3]int

	for i := range nums {
		var (
			left  = i + 1
			right = len(nums) - 1
		)

		for left < right {
			sum := nums[i] + nums[right] + nums[left]
			if sum == target {
				return target
			} else if sum > target {
				cSub := target - sum
				if cSub < 0 {
					cSub = cSub * -1
				}
				if cSub < minSub {
					minSub = cSub
					retArry = [3]int{nums[i], nums[left], nums[right]}
				}
				right -= 1
				continue
			} else {
				cSub := target - sum
				if cSub < 0 {
					cSub = cSub * -1
				}
				if cSub < minSub {
					minSub = cSub
					retArry = [3]int{nums[i], nums[left], nums[right]}
				}
				left += 1
				continue
			}
		}
	}
	return retArry[0] + retArry[1] + retArry[2]
}

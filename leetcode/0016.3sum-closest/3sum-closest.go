package problem0016

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ret, delta := 0, math.MaxInt64

	for i := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s > target:
				r--
				if s-target < delta {
					delta = s - target
					ret = s
				}
			case s < target:
				l++
				// 注意delta为正
				if target-s < delta {
					delta = target - s
					ret = s
				}
			default:
				return s
			}
		}
	}
	return ret
}

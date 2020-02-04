package problem0015

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	ret := [][]int{}
	if len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)
	if nums[0] > 0 {
		return ret
	}
	var (
		left   int
		right  int
		middle int
	)
	for i := range nums {
		if nums[i] > 0 {
			return ret
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left = i
		right = len(nums) - 1
		middle = left + 1
		for middle < right {
			cSum := nums[left] + nums[middle] + nums[right]
			if cSum == 0 {
				ret = append(ret, []int{nums[left], nums[middle], nums[right]})
				for nums[middle] == nums[middle+1] && middle+1 < right {
					middle += 1
				}
				for nums[right] == nums[right-1] && middle+1 < right {
					right -= 1
				}
				middle += 1
				right -= 1
			} else if cSum > 0 {
				right -= 1
			} else if cSum < 0 {
				middle += 1
			}
		}
	}
	return ret
}

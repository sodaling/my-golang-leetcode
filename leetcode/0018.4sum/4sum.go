package problem0018

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	for i := range nums {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, len(nums)-1
			for l < r {
				s := nums[i] + nums[j] + nums[l] + nums[r]
				switch {
				case s > target:
					r--
				case s < target:
					l++
				default:
					ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})
				loop:
					for l < r {
						switch {
						case nums[l+1] == nums[l]:
							l++
						case nums[r-1] == nums[r]:
							r--
						default:
							l++
							r--
							break loop
						}
					}
				}
			}
		}
	}
	return ret
}

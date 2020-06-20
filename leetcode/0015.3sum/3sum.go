package problem0015

import (
	"sort"
)

func threeSum1(nums []int) [][]int {
	// 两次循环
	// 通过了，但是贼慢，我估计是慢在判重的地方
	sort.Ints(nums)
	var ret [][]int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		need := 0 - nums[i]
		tempMap := make(map[int]int)
		for j := i + 1; j < len(nums); j++ {
			need2 := need - nums[j]
			if v, ok := tempMap[need2]; ok {
				ret = append(ret, []int{nums[i], nums[v], nums[j]})
			} else {
				tempMap[nums[j]] = j
			}
		}
	}
	return ret
}

// 去重不要用太复杂的方法，利用排序后的特性去重，然后利用双指针(排序后)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		// 避免添加重复的结果
		// i>0 是为了防止nums[i-1]溢出
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {

			case s > 0:
				r--
			case s < 0:
				l++
			default:
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
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
	return ret
}

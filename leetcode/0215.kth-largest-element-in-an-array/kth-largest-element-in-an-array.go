package problem0215

func findKthLargest(nums []int, k int) int {
	// 方法没错,可以不用递归优化
	size := len(nums)
	if size == 0 {
		return -1
	}
	return split(nums, k)
}

func split(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	temp, i := nums[0], 1
	left, right := 0, len(nums)-1
	for left < right {
		if nums[i] < temp {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		} else {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++
		}
	}
	nums[left] = temp
	if right+1 > k {
		return split(nums[:right], k)
	} else if right+1 < k {
		return split(nums[right+1:], k-right-1)
	}
	return nums[right]
}

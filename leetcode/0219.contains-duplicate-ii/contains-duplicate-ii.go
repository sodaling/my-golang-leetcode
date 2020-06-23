package problem0219

func containsNearbyDuplicate(nums []int, k int) bool {
	if k <= 0 {
		return false
	}

	// 利用 m 记录 a[i]中的值，每次出现的(索引号+1)
	// 因为从0开始，所以加1
	m := make(map[int]int, len(nums))

	for i, n := range nums {
		if m[n] != 0 && i-(m[n]-1) <= k {
			return true
		}
		m[n] = i + 1
	}

	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {
	// 利用滑动窗口
	if k <= 0 {
		return false
	}
	left, right := 0, 0
	window := make(map[int]int)
	for right < len(nums) {
		r := nums[right]
		right++
		if v := window[r]; v == 1 {
			return true
		} else {
			window[r]++
		}
		for right-left > k {
			l := nums[left]
			left++
			window[l]--
		}
	}
	return false
}

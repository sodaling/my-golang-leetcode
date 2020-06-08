package problem0167

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		sum := nums[right] + nums[left]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return nil
}

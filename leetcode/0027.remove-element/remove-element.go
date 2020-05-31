package problem0027

func removeElement(nums []int, val int) int {
	pre := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[pre], nums[i] = nums[i], nums[pre]
			pre++
		}
	}
	return pre
}

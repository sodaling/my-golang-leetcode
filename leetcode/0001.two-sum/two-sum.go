package problem0001

func twoSum(nums []int, target int) []int {
	tempMap := make(map[int]int, len(nums))
	for i := range nums {
		need := target - nums[i]
		if j, ok := tempMap[need]; !ok {
			tempMap[nums[i]] = i
		} else {
			return []int{i, j}
		}
	}
	return nil
}

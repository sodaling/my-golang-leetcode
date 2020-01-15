package _300_longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	numLength := len(nums)
	dp := make([]int, numLength+1)
	dp[0] = 0
	var retMax int
	for i := 1; i <= numLength; i++ {
		var iMax int
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] && dp[j] > iMax {
				iMax = dp[j]
			}
		}
		dp[i] = iMax + 1
		if dp[i] > retMax {
			retMax = dp[i]
		}
	}
	return retMax
}

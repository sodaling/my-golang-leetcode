package problem0003

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	left, right := 0, 0
	longest := 1
	strMap := make(map[byte]int)
	for right < len(s) {
		if _, ok := strMap[s[right]]; !ok {
			strMap[s[right]] = right
			if (right-left)+1 > longest {
				longest = right - left + 1
			}
		} else {
			for ; left < strMap[s[right]]+1; left++ {
				delete(strMap, s[left])
			}
			strMap[s[right]] = right
		}
		right++
	}
	return longest
}

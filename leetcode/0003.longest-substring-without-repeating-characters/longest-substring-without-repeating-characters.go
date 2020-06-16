package problem0003

func lengthOfLongestSubstring2(s string) int {
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

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	longest := 0
	left, right := 0, 0
	for right < len(s) {
		r := s[right]
		right++
		window[r]++
		for window[r] > 1 {
			l := s[left]
			left++
			if _, ok := window[l]; ok {
				window[l]--
			}
		}
		if right-left > longest {
			longest = right - left
		}
	}
	return longest
}

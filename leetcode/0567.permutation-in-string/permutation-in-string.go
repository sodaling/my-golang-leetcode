package problem0567

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}
	window := make(map[byte]int)
	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		r := s2[right]
		right++
		if _, ok := need[r]; ok {
			window[r]++
			if window[r] == need[r] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			l := s2[left]
			left++
			if _, ok := need[l]; ok {
				if need[l] == window[l] {
					valid--
				}
				window[l]--
			}
		}
	}
	return false
}

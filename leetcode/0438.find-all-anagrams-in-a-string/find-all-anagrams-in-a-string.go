package problem0438

func findAnagrams(s string, p string) []int {
	if len(s) == 0 {
		return nil
	}
	var ret []int
	pLen := len(p)
	strMap := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		strMap[p[i]] = -1
	}
	left, right := 0, 0

	for right < len(s) {
		if strMap[s[right]] == -1 {
			strMap[s[right]] = right
			if right-left+1 == pLen {
				ret = append(ret, left)
			}
		} else if _, ok := strMap[s[right]]; ok {
			for ; left < strMap[s[right]]+1; left++ {
				strMap[s[left]] = -1
			}
			if right-left+1 == pLen {
				ret = append(ret, left)
			}
			strMap[s[right]] = right
		} else {
			left = right + 1
			for key := range strMap {
				strMap[key] = -1
			}
		}
		right++
	}
	return ret
}

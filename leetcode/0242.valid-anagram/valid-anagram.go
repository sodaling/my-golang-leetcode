package problem0242

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sRune := []rune(s)
	tRune := []rune(t)
	tempMap := make(map[rune]int)
	for i := range sRune {
		tempMap[sRune[i]]++
		tempMap[tRune[i]]--
	}
	for key := range tempMap {
		if tempMap[key] != 0 {
			return false
		}
	}
	return true
}

func isAnagram(s string, t string) bool {
	// 利用golang数组可以比较的特性。只有对应元素相等时候是等的
	a := [26]int{}
	b := [26]int{}
	for _, v := range s {
		a['z'-v] += 1
	}
	for _, v := range t {
		b['z'-v] += 1
	}
	return a == b
}

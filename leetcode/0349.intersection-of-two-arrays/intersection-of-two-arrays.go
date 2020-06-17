package problem0349

func intersection2(a1, a2 []int) []int {
	tempMap := make(map[int]struct{})
	tempMap2 := make(map[int]struct{})
	for i := range a1 {
		tempMap[a1[i]] = struct{}{}
	}
	for i := range a2 {
		if _, ok := tempMap[a2[i]]; ok {
			tempMap2[a2[i]] = struct{}{}
		}
	}
	ret := make([]int, 0, len(tempMap2))
	for i := range tempMap2 {
		ret = append(ret, i)
	}
	return ret
}

func intersection(a1, a2 []int) []int {
	res := []int{}
	m1 := getInts(a1)
	m2 := getInts(a2)

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}
	// m1 是较短的字典，会快一些
	for n := range m1 {
		if m2[n] {
			res = append(res, n)
		}
	}

	return res
}

// 清理重复的值，也便于查询
func getInts(a []int) map[int]bool {
	res := make(map[int]bool, len(a))

	for i := range a {
		res[a[i]] = true
	}

	return res
}

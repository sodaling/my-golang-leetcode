package problem0350

func intersect(a1, a2 []int) []int {
	tempMap := make(map[int]int)
	tempMap2 := make(map[int]int)
	for i := range a1 {
		tempMap[a1[i]]++
	}
	length := 0
	for i := range a2 {
		if _, ok := tempMap[a2[i]]; ok {
			if tempMap2[a2[i]] < tempMap[a2[i]] {
				tempMap2[a2[i]]++
				length++
			}
		}
	}
	ret := make([]int, 0, length)
	for i := range tempMap2 {
		for j := 0; j < tempMap2[i]; j++ {
			ret = append(ret, i)
		}
	}
	return ret

}

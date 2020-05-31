package problem0026

func removeDuplicates(a []int) int {
	if len(a) == 0 {
		return 0
	}
	pre := 1
	for i := 1; i < len(a); i++ {
		if a[i] != a[pre-1] {
			a[i], a[pre] = a[pre], a[i]
			pre++
		}
	}
	return pre
}

package problem0454

func fourSumCount(A []int, B []int, C []int, D []int) int {
	CAndDMap := make(map[int]int)
	ret := 0
	for i := range C {
		for j := range D {
			CAndDMap[C[i]+D[j]]++
		}
	}
	for i := range A {
		for j := range B {
			if v, ok := CAndDMap[0-A[i]-B[j]]; ok {
				ret = ret + v
			}
		}
	}
	return ret

}

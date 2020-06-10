package problem0209

func minSubArrayLen(s int, a []int) int {
	left, right := 0, -1
	ret := len(a) + 1
	sumNow := 0
	for left < len(a) {
		if sumNow < s && right+1 < len(a) {
			right++
			sumNow += a[right]
		} else if sumNow >= s {
			ret = min(right-left+1, ret)
			sumNow -= a[left]
			left++
		} else {
			break
		}
	}
	if ret == len(a)+1 {
		return 0
	} else {
		return ret
	}
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

package problem0080

func removeDuplicates(a []int) int {
	// 只有一种情况下不用换,就是和前面两个都相等了,那就直接i+2,也就是continue,反正只要是换了,pre就得加1,前面的加1后前面的正确性得保证
	// 这道题双100...
	size := len(a)
	if size <= 2 {
		return size
	}
	pre := 2
	for i := 2; i < len(a); i++ {
		if a[i] == a[pre-2] {
			continue
		}
		a[i], a[pre] = a[pre], a[i]
		pre++
	}
	return pre
}

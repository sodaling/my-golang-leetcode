package problem0088

// 本题的要求是，把nums1的前m项和nums2的前n项合并，放入nums1中。
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 三个指针,从数组后面,也就是从到小进行归并.值得注意的是,因为k指针肯定追不上m指针,所以不会清洗掉nums1的数据
	m, n, k := m-1, n-1, m+n-1
	for k >= 0 {
		if m < 0 {
			nums1[k] = nums2[n]
			n--
			k--
			continue
		}
		if n < 0 {
			nums1[k] = nums1[m]
			m--
			k--
			continue
		}
		if nums1[m] >= nums2[n] {
			nums1[k] = nums1[m]
			m--
		} else {
			nums1[k] = nums2[n]
			n--
		}
		k--
	}
}

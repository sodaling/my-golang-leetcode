package main

import "fmt"

func HeapSort(array []int) {

	// 首先构造大根堆
	m := len(array)
	s := m / 2
	for i := s; i >= 0; i-- {
		// 注意是--
		heap(array, i, m-1)
	}
	for i := m - 1; i > 0; i-- {
		// 最大的放到最后,然后开始重新构造堆
		array[i], array[0] = array[0], array[i]
		heap(array, 0, i-1)
	}

}

func heap(array []int, i, end int) {
	// i是最后一个叶子节点的父节点
	l := i*2 + 1
	if l > end {
		return
	}
	n := l
	r := 2*i + 2
	if r <= end && array[r] > array[l] {
		n = r
	}
	if array[i] > array[n] {
		// 这种情况下就没必要替换了，直接return可以了
		return
	}
	array[n], array[i] = array[i], array[n]
	heap(array, n, end)
}

func main() {
	array := []int{
		55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9,
	}
	HeapSort(array)
	fmt.Println(array)
}

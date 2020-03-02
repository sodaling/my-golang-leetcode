package main

import "fmt"

func MergeSort(arr []int, l, r int) {
	if r <= l {
		return
	}
	mid := (r + l) / 2
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func merge(arr []int, l int, m int, r int) {
	temp := make([]int, r-l+1)
	copy(temp, arr[l:r+1])
	left := 0
	right := m - l + 1
	for i := l; i <= r; {
		if temp[left] >= temp[right] {
			arr[i] = temp[right]
			right++
			i++
			if right == len(temp) {
				for left < m-l+1 && i <= r {
					arr[i] = temp[left]
					left++
					i++
				}
			}
		} else {
			arr[i] = temp[left]
			left++
			i++
			if left == m-l+1 {
				for right <= len(temp)-1 && i <= r {
					arr[i] = temp[right]
					right++
					i++
				}
			}
		}
	}
}

func main() {
	a := []int{6, 2, 3, 1, 5, 412, 31, 142342, 12, 11, 31, 42, 53, 6, 2}
	MergeSort(a, 0, len(a)-1)
	fmt.Println(a)
}

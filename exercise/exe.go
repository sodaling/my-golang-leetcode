package main

import (
	"fmt"
	"reflect"
)

func Quick2Sort(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	left, right := 0, len(values)
	for i < right {
		if values[i] == mid {
			i++
		} else if values[i] > mid {
			right--
			values[right], values[i] = values[i], values[right]
		} else {
			left++
			values[i], values[left] = values[left], values[i]
			i++
		}
	}
	values[0], values[left] = values[left], values[0]
	Quick2Sort(values[:left+1])
	Quick2Sort(values[right:])
}

func heapSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	buildMaxHeap(arr, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		swap(arr, 0, i)
		arrLen -= 1
		heapify(arr, 0, arrLen)
	}
}

func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func heapify(arr []int, i, arrLen int) {
	left := i*2 + 1
	right := i*2 + 2
	largest := i
	if left < arrLen && arr[left] > arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, largest, i)
		heapify(arr, largest, arrLen)
	}
}

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
	str := "Hello,世界"
	//utf-8遍历
	for i := 0; i < len(str); i++ {
		ch := str[i]
		ctype := reflect.TypeOf(ch)
		fmt.Printf("%s ", ctype)
	}
	fmt.Println("=============>Unicode遍历")
	//Unicode遍历
	for _, ch1 := range str {
		ctype := reflect.TypeOf(ch1)
		fmt.Printf("%s ", ctype)
		fmt.Printf("%c", ch1)
	}
}

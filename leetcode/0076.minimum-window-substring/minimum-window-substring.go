package problem0076

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	valid := 0
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := 0, 0
	retLeft, retLen := 0, math.MaxInt64
	for right < len(s) {
		r := s[right]
		right++
		if _, ok := need[r]; ok {
			window[r]++
			if need[r] == window[r] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left < retLen {
				retLen = right - left
				retLeft = left
			}
			l := s[left]
			left++
			if _, ok := need[l]; ok {
				if window[l] == need[l] {
					valid--
				}
				window[l]--
			}
		}
	}

	if retLen == math.MaxInt64{
		return ""
	} else {
		fmt.Println(retLeft, retLen)
		return s[retLeft : retLeft+retLen]
	}
}

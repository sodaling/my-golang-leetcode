package problem0125

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	head, tail := 0, len(s)-1
	for head < tail {
		for head < tail && !isChar(s[head]) {
			head++
		}
		for head < tail && !isChar(s[tail]) {
			tail--
		}
		if s[head] != s[tail] {
			return false
		}
		head++
		tail--
	}
	return true

}

// 判断 c 是否是需要检查的字符
func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}

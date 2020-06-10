package problem0344

func reverseString(s string) string {
	sRune := []rune(s)
	head, tail := 0, len(sRune)-1
	for head < tail {
		sRune[head], sRune[tail] = sRune[tail], sRune[head]
		head++
		tail--
	}
	return string(sRune)
}

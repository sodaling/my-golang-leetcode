package problem0438

func findAnagramsErr(s string, p string) []int {
	// 下面的是不行的，因为一旦p有重复字符的话就得不到想要的结果
	if len(s) == 0 {
		return nil
	}
	var ret []int
	pLen := len(p)
	strMap := make(map[byte]struct{})
	arrMap := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		strMap[p[i]] = struct{}{}
	}
	left, right := -1, 0

	for ; right < len(s); right++ {
		if _, ok := strMap[s[right]]; ok {
			arrLen := len(arrMap)
			if arrLen == 0 {
				arrMap[s[right]], left = right, right
			} else if _, y := arrMap[s[right]]; y {
				for ; left < arrMap[s[right]]+1; left++ {
					delete(arrMap, s[left])
				}
				arrMap[s[right]] = right
				if len(arrMap) == pLen {
					ret = append(ret, left)
				}
			} else {
				arrMap[s[right]] = right
				if len(arrMap) == pLen {
					ret = append(ret, left)
				}
			}
		} else {
			left = right + 1
			arrMap = make(map[byte]int)
		}
	}
	return ret
}
func findAnagrams(s string, p string) []int {
	need := make(map[byte]int)
	window:= make(map[byte]int)
	for i:=0;i<len(p);i++{
		need[p[i]]++
	}
	valid := 0
	var ret []int
	left,right := 0,0
	for right<len(s){
		// 窗口右移
		c := s[right]
		right++
		if _,ok:=need[c];ok{
			window[c]++
			if need[c]==window[c]{
				valid++
			}
		}
		// 窗口右移
		for right-left>=len(p){
			// 满足条件则更新返回数组
			if valid==len(need){
				ret = append(ret,left)
			}
			c = s[left]
			left++
			if _,ok:=need[c];ok{
				if window[c]==need[c]{
					valid--
				}
				window[c]--
			}
		}
	}
	return ret
}

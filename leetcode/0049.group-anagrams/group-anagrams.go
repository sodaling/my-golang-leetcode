package problem0049

import (
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortRunes) Len() int {
	return len(s)
}
func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
func groupAnagrams(ss []string) [][]string {
	//利用排序
	tempMap := make(map[string][]int)
	var ret [][]string
	for i := range ss {
		sortString := SortString(ss[i])
		tempMap[sortString] = append(tempMap[sortString], i)
	}
	for _, v := range tempMap {
		tempSlice := make([]string, len(v))
		for i, index := range v {
			tempSlice[i] = ss[index]
		}
		ret = append(ret, tempSlice)
	}
	return ret
}

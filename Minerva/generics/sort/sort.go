package main

import "sort"

func main() {
	_ = make(map[string][]string)
}

func sortString(a string) string {
	s := []byte(a)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return string(s)
}

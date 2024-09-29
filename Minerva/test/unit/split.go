package unit

import (
	"strings"
)

func split(s, sep string) []string {
	count := strings.Count(s, sep)
	res := make([]string, 0, count+1)
	index := strings.Index(s, sep)
	for index >= 0 {
		res = append(res, s[:index])
		s = s[index+len(sep):]
		index = strings.Index(s, sep)
	}
	res = append(res, s)
	return res
}

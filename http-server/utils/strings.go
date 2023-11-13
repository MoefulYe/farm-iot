package utils

import "strings"

// FirstToUpper 让开头字母大写
func FirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	return strings.ToUpper(string(str[0])) + str[1:]
}

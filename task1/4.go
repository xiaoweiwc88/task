package main

import (
	"fmt"
)

func longestCommanPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	p := strs[0]
	for _, s := range strs[1:] {
		i := 0
		for i < len(p) && i < len(s) && p[i] == s[i] {
			i++
		}
		p = p[:i]
		if p == "" {
			return ""
		}
	}
	return p
}

func main() {
	str := []string{"flower", "flow", "flight"}
	str1 := []string{"dog","racecar","car"}


	fmt.Printf("%s is longestCommonPrefix\n", longestCommanPrefix(str))
	fmt.Printf("%s is longestCommonPrefix\n", longestCommanPrefix(str1))
}

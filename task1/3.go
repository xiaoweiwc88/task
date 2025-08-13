package main

import "fmt"

func isVlidMap(s string) bool {
	n := len(s)
	if n % 2 == 1 {
		return false
	}
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else if c == ')' || c == '}' || c == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	arr := []string{"()", "()[]{}", "(]", "([])", "([)]"}
	for _, v := range arr {
		if isVlidMap(v) {
			fmt.Printf("%s is vallid\n", v)
		} else {
			fmt.Printf("%s is invallid\n", v)
		}
	}
	return
}

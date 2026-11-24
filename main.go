package main

import "fmt"

// IsBalanced reports whether every bracket in s — (), [], {} — closes the
// most recently opened, unclosed bracket of its own kind.
func IsBalanced(s string) bool {
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(IsBalanced("{[()]}"), IsBalanced("(]"))
}

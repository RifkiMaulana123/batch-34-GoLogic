package main

import "fmt"

func main() {
	fmt.Println(checkBraces("(())"))
	fmt.Println(checkBraces("()()"))
	fmt.Println(checkBraces("((()"))
	fmt.Println(checkBraces("(()))((())"))
}

func checkBraces(s string) bool {
	count := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			count--
		}
		if count < 0 {
			return false
		}
	}

	return count == 0
}

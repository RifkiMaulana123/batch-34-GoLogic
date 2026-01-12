package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("Kasur ini rusak"))
	fmt.Println(isPalindrome("tamaT"))
	fmt.Println(isPalindrome("Aku Usa"))
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i]|32 != s[len(s)-1-i]|32 {
			return false
		}
	}

	return true
}

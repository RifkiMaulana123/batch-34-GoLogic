package main

import "fmt"

func main() {
	fmt.Println(isNumberPalindrome(121))
	fmt.Println(isNumberPalindrome(2147447412))
	fmt.Println(isNumberPalindrome(333))
	fmt.Println(isNumberPalindrome(110))
	fmt.Println(isNumberPalindrome(11))

}

func isNumberPalindrome(n int) bool {
	number := n
	reverse := 0

	for n > 0 {
		reverse = reverse*10 + n%10
		n /= 10
	}
	return number == reverse
}

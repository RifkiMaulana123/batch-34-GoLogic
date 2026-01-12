package main

import "fmt"

func main() {
	extractDigit(12234)
	extractDigit(5432)
	extractDigit(1278)
}

func extractDigit(n int) {
	for n > 0 {
		fmt.Print(n%10, " ")
		n = n / 10
	}
	fmt.Println()

}

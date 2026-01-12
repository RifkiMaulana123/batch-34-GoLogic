package main

import "fmt"

func main() {
	pyramid()
}

func pyramid() {
	var n int
	fmt.Print("Input jumlah baris pyramid : ")
	fmt.Scan(&n)

	for i := n; i >= 1; i-- {

		for j := i; j >= 1; j-- {
			fmt.Print(j, " ")
		}

		for j := 2; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

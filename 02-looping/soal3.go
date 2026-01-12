package main

import "fmt"

func main() {
	triangle()
}

func triangle() {

	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j >= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print(" ")

		for j := 1; j <= 5; j++ {
			if i+j >= 6 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

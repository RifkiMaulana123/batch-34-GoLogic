package main

import "fmt"

func main() {
	deret(9)
	fmt.Println()
	deret(5)
}

func deret(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {

			if j%2 == 1 {
				fmt.Print(i)
			} else {
				fmt.Print(n - i + 1)
			}

		}
		fmt.Println()
	}

}

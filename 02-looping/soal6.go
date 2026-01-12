package main

import "fmt"

func main() {
	pola(9)
	fmt.Println()
	pola(5)
}

func pola(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {

			if i%2 == 1 {
				if j%2 == 0 {
					fmt.Print(j)
				} else {
					fmt.Print("-")
				}
			} else {
				if j%2 == 1 {
					fmt.Print(j)
				} else {
					fmt.Print("-")
				}
			}
		}
		fmt.Println()
	}

}

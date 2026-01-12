package main

import "fmt"

func main() {
	fmt.Println(reverse("DCBA"))
	fmt.Println(reverse("tamaT"))
	fmt.Println(reverse("XYnb"))

}

func reverse(s string) string {
	result := ""

	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}

	return result
}

package main

import "fmt"

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func main() {
	a := "sula"
	b := reverse("alus")
	if a == b {
		fmt.Println(a, "polindrom k", reverse(b))
		fmt.Println(a, b, "=", true)
	} else {
		fmt.Println(a, "polindrom k", reverse(b))
		fmt.Println(a, b, "=", false)
	}
}

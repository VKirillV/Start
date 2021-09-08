package main

import "fmt"

func main() {
	fizz := "fizz"
	buss := "buss"
	for x := 1; x <= 100; x++ {
		if x%3 == 0 && x%5 == 0 {
			fmt.Println(fizz, buss)
		} else if x%3 == 0 {
			fmt.Println(fizz)
		} else if x%5 == 0 {
			fmt.Println(buss)
		} else {
			fmt.Println(x)
		}
	}
}

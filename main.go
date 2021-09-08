package main

import "fmt"

func main() {
	a := [5]string{"Hello", "World", "in", "a", "frame"}
	for x, b := range a {
		fmt.Println(x, b)
	}
}

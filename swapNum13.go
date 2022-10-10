package main

import "fmt"

// swap значений местами, без третьей переменной

func main() {
	a, b := 2, 4
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}

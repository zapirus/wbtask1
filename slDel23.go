package main

import "fmt"

//Удалить i-ый элемент из слайса.

func delSl(x []int, xs int) []int {
	x = append(x[:xs], x[xs+1:]...)
	return x
}

func main() {
	sl := []int{1, 23, 4, 5, 67, 8}
	var s int
	fmt.Scan(&s)
	fmt.Println(delSl(sl, s))
}

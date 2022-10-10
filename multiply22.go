package main

import (
	"fmt"
	"math"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
//значение которых > 2^20.

func Checking(x int) bool {
	if x > int(math.Pow(2, 20)) {
		return true
	}
	return false
}

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	if Checking(a) || Checking(b) {
		fmt.Printf("%d, %d, %d, %d", a*b, a/b, a+b, a-b)
	} else {
		fmt.Println("Значения малы")
	}

}

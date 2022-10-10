package main

import (
	"fmt"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0

func Bit(i int64, x int64) string {

	return fmt.Sprintf("%b, %d", x|1<<i, x|1<<i)
}

func main() {
	var x int64 = 30
	fmt.Println(Bit(7, x))
}

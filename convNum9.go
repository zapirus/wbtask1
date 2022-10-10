package main

import (
	"fmt"
)

// Разработать конвейер чисел.
//Даны два канала:
//в первый пишутся числа (x) из массива, во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

func writing(ch1 chan int, mass []int) {

	for _, i2 := range mass {
		ch1 <- i2
	}
	close(ch1)

}

func wrMultiply(ch2 chan int, mass []int) {

	for _, i2 := range mass {
		ch2 <- i2 * 2
	}
	close(ch2)

}

func main() {
	massive := []int{1, 2, 3, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go writing(ch1, massive)
	go wrMultiply(ch2, massive)

	for i2 := range ch2 {
		fmt.Printf("%d ", i2)
	}
}

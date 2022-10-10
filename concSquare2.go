package main

import (
	"fmt"
	"sync"
)

//конкурентная функция, которая рассчитывет квадрат с массива)

var wg sync.WaitGroup
var ms = [5]int{2, 4, 6, 8, 10}

func concSquare() {
	wg.Add(1)
	go func(mass [5]int) {
		defer wg.Done()
		for _, i2 := range ms {
			fmt.Println(i2 * i2)
		}
	}(mass)
}

func main() {

	concSquare()
	wg.Wait()
}

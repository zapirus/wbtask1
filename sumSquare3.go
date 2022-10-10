package main

import (
	"fmt"
	"sync"
)

//конкурентная функция, которая рассчитывет квадрат с массива)

var wgsq sync.WaitGroup
var mass = [5]int{2, 4, 6, 8, 10}

func smSq() {
	wgsq.Add(1)
	go func(mass [5]int) {
		sum := 0
		defer wgsq.Done()
		for _, i2 := range mass {
			sum += i2 * i2
		}

		fmt.Println(sum)
	}(mass)
}
func main() {
	smSq()
	wgsq.Wait()

}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.

func write(channels chan int) {
	go func() {
		for {
			rd := rand.Int()
			channels <- rd
		}
	}()
}

func read(channels chan int) {
	go func() {
		for ch := range channels {
			fmt.Println(ch)
		}
	}()
}

func main() {
	c := make(chan int)

	write(c)
	read(c)

	select {
	case <-time.After(7 * time.Second):
		fmt.Println("Exit")
	}
}

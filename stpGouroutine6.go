package main

import (
	"fmt"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

func main() {

	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "запись в канал"
		defer close(ch)
	}()
	select {
	case out := <-ch:
		fmt.Println(out)
	case <-time.After(4 * time.Second):
		fmt.Println("тайм-аут")
	}
}

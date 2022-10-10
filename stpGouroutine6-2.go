package main

import (
	"fmt"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	chs := make(chan int)
	go func() {
		defer close(chs)
		time.Sleep(4 * time.Second)
		w.Wait()
	}()

	select {
	case <-chs:
		return false
	case <-time.After(t):
		return true

	}
}
func main() {

	var w sync.WaitGroup
	w.Add(1)
	var t int
	fmt.Scan(&t)

	dr := time.Duration(t) * time.Second
	fmt.Printf("Период ожидания %s\n", dr)

	if timeout(&w, dr) {
		fmt.Println("тайм-аут1")
	} else {
		fmt.Println("ok")
	}

	w.Done()

	if timeout(&w, dr) {
		fmt.Println("тайм-аут2")
	} else {
		fmt.Println("ok")
	}

}

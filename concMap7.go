package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

var ws sync.WaitGroup

func main() {

	mp := make(map[string]int)
	mutex := sync.Mutex{}

	for i := 0; i < 20; i++ {
		ws.Add(1)
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			mp["ct"]++
			ws.Done()
		}()

	}
	ws.Wait()

	fmt.Println(mp)

}

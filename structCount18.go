package main

import (
	"fmt"
	"sync"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

type Structs struct {
	X     int
	mutex sync.Mutex
}

func (s *Structs) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	s.mutex.Lock()
	s.X++
	defer s.mutex.Unlock()

}

func main() {
	s := 0
	wgs := &sync.WaitGroup{}
	inc := &Structs{mutex: sync.Mutex{}}
	for s < 500 {
		wgs.Add(1)
		go inc.inc(wgs)
		s++
	}
	wgs.Wait()
	fmt.Println("Increment ", inc.X)
}

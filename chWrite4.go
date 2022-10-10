package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.

type Client struct {
	id      int
	integer int
}

type Data struct {
	job    Client
	square int
}

var (
	size    = 1
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

func worker(w *sync.WaitGroup) {
	for c := range clients {
		square := c.integer * c.integer
		output := Data{c, square}

		data <- output
		time.Sleep(time.Second)
	}
	w.Done()
}

func makeWP(n int) {
	var w sync.WaitGroup
	for i := 0; i < n; i++ {
		w.Add(1)
		go worker(&w)

	}
	w.Wait()
	close(data)
}

func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}

func main() {
	fmt.Println("Capacity of clients:", cap(clients))
	fmt.Println("Capacity of data:", cap(data))

	nJobs := 9223372036854775807

	nWorkers := 10

	go create(nJobs)
	//finished := make(chan interface{})
	for {

		go func() {
			for d := range data {
				fmt.Printf("Client ID: %d\tint: ", d.job.id)
				fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
			}
			//finished <- true
		}()

		makeWP(nWorkers)
		//fmt.Printf(": %v\n", <-finished)

	}

	//Реализовать постоянную запись данных в канал (главный поток).
	//Реализовать набор из N воркеров,
	//которые читают произвольные данные из канала и выводят в stdout.
	//Необходима возможность выбора количества воркеров при старте.
}

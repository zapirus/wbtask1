package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep

func Sleep(d time.Duration) {
	timer := time.NewTimer(d)
	<-timer.C
}

func main() {
	Sleep(2 * time.Second)
	fmt.Println("The end")
}

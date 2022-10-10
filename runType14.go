package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func runType(a interface{}) {
	st := reflect.TypeOf(a).Kind()
	ints := reflect.TypeOf(a).Kind()
	bl := reflect.TypeOf(a).Kind()
	ch := reflect.TypeOf(a).Kind()

	if st == reflect.String {
		fmt.Println("string")
	} else if ints == reflect.Int {
		fmt.Println("int")
	} else if bl == reflect.Bool {
		fmt.Println("bool")
	} else if ch == reflect.Chan {
		fmt.Printf("chan")
	} else {
		fmt.Println("404")
	}
}

func main() {

	runType(1)

}

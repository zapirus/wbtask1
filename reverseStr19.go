package main

import "fmt"

//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
//Символы могут быть unicode.
func reverse(str string) (result string) {
	for _, s := range str {
		result = string(s) + result
	}
	return result
}

func main() {

	var str string
	fmt.Scan(&str)

	strReverse := reverse(str)
	fmt.Println(str)
	fmt.Println(strReverse)
}

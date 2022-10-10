package main

import (
	"fmt"
	"strings"
)

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
func main() {
	strToSlice := strings.Split("cat, cat, dog, cat, tree", " ")
	mp := make(map[int]string)

	for i, r := range strToSlice {
		mp[i] = r
	}
	fmt.Println(mp)

}

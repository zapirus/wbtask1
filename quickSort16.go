package main

import (
	"fmt"
	"math/rand"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	l, r := 0, len(a)-1

	ct := rand.Int() % len(a)

	a[ct], a[r] = a[r], a[ct]
	for i, _ := range a {
		if a[i] < a[r] {
			a[l], a[i] = a[i], a[l]
			l++
		}
	}
	a[l], a[r] = a[r], a[l]

	quicksort(a[:l])
	quicksort(a[l+1:])
	return a

}

func main() {
	nums := []int{4, 0, 3, 7, 1, 3, 10, 15, 12, 6}
	fmt.Println(nums)
	fmt.Println(quicksort(nums))

}

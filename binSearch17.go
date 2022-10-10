package main

import "fmt"

//Реализовать бинарный поиск встроенными методами языка.

func binSearch(a []int, x int) int {
	result := 0
	st, end := 0, len(a)
	for st <= end {

		mid := (st + end) / 2
		if a[mid] == x {
			result = mid
			break
		} else if a[mid] < x {
			st = mid + 1
		} else if a[mid] > x {
			end = mid - 1
		}
	}
	return result
}

func main() {

	nums := []int{4, 0, 3, 7, 12, 2134, 3, 10, 15, 12, 6}
	fmt.Println(binSearch(nums, 4))
}

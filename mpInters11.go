package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func mpInters(a, b []string) []string {
	m := make(map[string]int)
	mpIntr := make([]string, 0)
	for _, v := range a {
		m[v] += 1
	}

	for _, v := range b {
		m[v] += 1
	}

	for k := range m {
		if m[k] > 1 {
			mpIntr = append(mpIntr, k)
		}
	}
	return mpIntr
}

func main() {
	a := []string{"a", "b", "c", "d", "e", "g"}
	b := []string{"a", "c", "f", "g", "t"}
	fmt.Println(mpInters(a, b))
}

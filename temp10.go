package main

import "fmt"

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

func main() {
	temp := []float32{100.3, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 33.8, 5.9, 1}
	mp := make(map[int][]float32)
	for _, n := range temp {
		j := n / 10
		mp[int(j)*10] = append(mp[int(j)*10], n)
	}

	fmt.Println(mp)
}

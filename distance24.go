package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры
//Point с инкапсулированными параметрами x,y и конструктором

type Point struct {
	x int
	y int
}

func AB(A, B Point) float64 {
	return math.Ceil(math.Sqrt(math.Pow(math.Abs(float64(B.x)-float64(A.x)), 2) + math.Pow(math.Abs(float64(B.y)-float64(A.y)), 2)))
}

func main() {
	p1 := Point{x: 10, y: 10}
	p2 := Point{x: 23, y: 30}
	fmt.Print(AB(p1, p2))
}

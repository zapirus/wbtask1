package main

import "fmt"

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
func unique(str string) (string, bool) {
	rn := []rune(str)
	ln := len(rn)
	for i := 0; i < ln; i++ {
		for j := i + 1; j < ln; j++ {
			if rn[j] == rn[i] {
				return str, false
			}
		}
	}
	return str, true
}

func main() {
	v := ""
	fmt.Scan(&v)

	fmt.Println(unique(v))
}

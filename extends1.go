package main

import "fmt"

//пример "нааследования" - встраивания

type Human struct {
	name string
}

func (h *Human) hello() {
	fmt.Println("hello my dear " + h.name)
}

func (b *Human) bye() string {
	return "bye"
}

type Action struct {
	Human
	Age string
}

func main() {
	act := Action{Human{name: "Ali"}, "31"}
	act2 := Action{Human{name: "Ruslan"}, "22"}

	fmt.Println("Info:"+" "+act.Human.name, act.Age)
	fmt.Println("Info:"+" "+act2.Human.name, act2.Age)
	act.hello()

	fmt.Println(act.bye())

}

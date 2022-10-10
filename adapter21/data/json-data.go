package data

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

type JsonDoc struct {
}

func (doc JsonDoc) ConvertToXml() string {
	return "xml"
}

type JsonDocAdapter struct {
	JsonDoc *JsonDoc
}

func (ar JsonDocAdapter) SendXml() {
	ar.JsonDoc.ConvertToXml()
	fmt.Println("Отправка данных")
}

package data_srv

import "fmt"

type DataService interface {
	SendXml()
}

type XmlDoc struct {
}

func (doc XmlDoc) SendXml() {
	fmt.Println("Отправка данных ")
}


package latihan

import "fmt"

type Person struct {
	name string
}

func Closure() {
	var orang = []*Person{
		{name: "Adi Kusdianto"},
		{name: "Bambang Wijaya"},
		{name: "Indah Lestari"},
		{name: "Dini Indriyani"},
		{name: "Renda Cokro"},
		{name: "Jajang Mulyana"},
		{name: "Bagaskara"},
		{name: "Mawar Indah"},
		{name: "Leila Nuryani"},
		{name: "Handoko Kusuma"},
	}

	contohClosure := func(persons []*Person) {

		for _, v := range persons {
			fmt.Println(v.name)
		}

	}

	contohClosure(orang)

}

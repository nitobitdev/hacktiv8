package latihan

import "fmt"

func Looping() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i+1, "Ganjil")
		} else {
			fmt.Println(i+1, "Genap")
		}
	}
}

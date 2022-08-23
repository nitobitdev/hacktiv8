package latihan

import "fmt"

func LoopingSlice() {
	var players = []string{
		"James",
		"Randy",
		"Dilan",
		"Lamelo",
		"Thohir",
		"Mars",
		"Antono",
		"Jamila",
		"Tison",
		"Azril",
	}

	for _, player := range players {
		fmt.Println(player)
	}
}

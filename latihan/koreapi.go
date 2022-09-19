package latihan

import (
	"fmt"
	"math/rand"
	"time"
)

type Korek struct {
	hits           int
	pemainTerakhir string
}

const BreakPoint = 11

func main() {
	korek := make(chan *Korek)

	selesai := make(chan *Korek)

	pemain := []string{"Pemain 1", "Pemain 2", "Pemain 3", "Pemain 4"}
	for _, p := range pemain {
		go bermain(p, korek, selesai)
	}
	korek <- new(Korek)

	berakhir(selesai)
}

func bermain(nama string, korek, selesai chan *Korek) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 100
	for {
		select {
		case k := <-korek:
			nilai := rand.Intn(max-min) + min
			time.Sleep(500 * time.Millisecond)
			k.hits++
			k.pemainTerakhir = nama
			fmt.Println("korek ada di", k.pemainTerakhir, "pada hit ke", k.hits, "dan mempunyai nilai", nilai)

			if nilai%BreakPoint == 0 {
				selesai <- k
				return
			}
			korek <- k
		}
	}
}

func berakhir(selesai chan *Korek) {
	for {
		select {
		case s := <-selesai:
			fmt.Println(s.pemainTerakhir, "kalah pada hit ke", s.hits)
			return
		}
	}
}

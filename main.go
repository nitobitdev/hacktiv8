package main

import (
	"fmt"
	"hacktiv8/latihan"
	"sync"
)

func main() {
	var db []*latihan.User
	user := latihan.NewUserService(db)
	names := []string{"Ucup", "Wahyono", "Jamilah", "Simorang", "Cahyo", "Andin", "Clara"}

	fmt.Println("________ WITHOUT GOROUTINE ________")
	// WITHOUT GOROUTINE
	for _, n := range names {
		res := user.Register(&latihan.User{Nama: n})
		fmt.Println(res)
	}

	getUser := user.GetUser()
	fmt.Println("________ GET USER WITHOUT GOROUTINE ________")

	for _, v := range getUser {
		fmt.Println(v.Nama)
	}

	fmt.Println("________ GOROUTINE ________")

	// WITH GOROUTINE
	var wg1 sync.WaitGroup
	wg1.Add(len(names))
	for _, n := range names {
		go func(name string) {
			res := user.Register(&latihan.User{Nama: name})
			fmt.Println(res)
			wg1.Done()
		}(n)
	}
	wg1.Wait()

	resGet := user.GetUser()
	fmt.Println("________ GET USER WITH GOROUTINE ________")
	var wg sync.WaitGroup
	wg.Add(len(resGet))
	for _, v := range resGet {
		go cetakNama(&wg, v.Nama)
	}
	wg.Wait()

}

func cetakNama(wg *sync.WaitGroup, nama string) {
	fmt.Println(nama)
	wg.Done()
}

package main

import (
	"fmt"
	"hacktiv8/latihan"
	"sync"
)

func main() {
	var db []*latihan.User
	userSvc := latihan.NewUserService(db)
	names := []string{"Ucup", "Wahyono", "Jamilah", "Simorang", "Cahyo", "Andin", "Clara"}
	var wg1 sync.WaitGroup
	wg1.Add(len(names))
	for _, n := range names {
		go func(name string) {
			res := userSvc.Register(&latihan.User{Nama: name})
			fmt.Println(res)
			wg1.Done()
		}(n)
	}
	wg1.Wait()

	resGet := userSvc.GetUser()
	fmt.Println("________ GET USER ________")
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

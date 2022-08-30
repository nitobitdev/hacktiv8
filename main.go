package main

import (
	"fmt"
	"hacktiv8/latihan"
)

func main() {
	var db []*latihan.User
	user := latihan.NewUserService(db)
	names := []string{"Ucup", "Wahyono", "Jamilah", "Simorang", "Cahyo", "Andin", "Clara"}
	for _, n := range names {
		res := user.Register(&latihan.User{Nama: n})
		fmt.Println(res)
	}

	getUser := user.GetUser()
	fmt.Println("________ GET USER ________")

	for _, v := range getUser {
		fmt.Println(v.Nama)
	}

	// var wg1 sync.WaitGroup
	// wg1.Add(len(names))
	// for _, n := range names {
	// 	go func(name string) {
	// 		res := userSvc.Register(&latihan.User{Nama: name})
	// 		fmt.Println(res)
	// 		wg1.Done()
	// 	}(n)
	// }
	// wg1.Wait()

	// resGet := userSvc.GetUser()
	// fmt.Println("________ GET USER ________")
	// var wg sync.WaitGroup
	// wg.Add(len(resGet))
	// for _, v := range resGet {
	// 	go cetakNama(&wg, v.Nama)
	// }
	// wg.Wait()

}

// func cetakNama(wg *sync.WaitGroup, nama string) {
// 	fmt.Println(nama)
// 	wg.Done()
// }

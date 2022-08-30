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

}

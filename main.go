package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string
}

var users = []User{
	{Name: "Budi Gunawan"},
	{Name: "Wawan Susanto"},
	{Name: "Jajan Nurjaman"},
}

var PORT = ":8080"

func main() {
	http.HandleFunc("/user", getUsers)
	http.HandleFunc("/register", createUser)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(users)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")

		newUser := User{
			Name: name,
		}

		users = append(users, newUser)

		fmt.Fprintf(w, name+" berhasil didaftarkan")
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)

}

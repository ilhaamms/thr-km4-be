package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	kolom := strings.Split(data, ",")

	age, err := strconv.Atoi(kolom[1])
	if err != nil {
		return User{}
	}

	return User{Name: kolom[0], Age: age, Address: kolom[2]}
}

func main() {
	user := ConvertData("Budi,30,Bandung")
	fmt.Println(user)
}

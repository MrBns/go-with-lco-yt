package main

import "fmt"

func main() {
	nazmul := User{Name: "H. Nazmul Hassan", Email: "mail@hnazmul.me", Status: true, Age: 21}

	fmt.Printf("%+v", nazmul.GetFullName())
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (p User) GetFullName() string {
	return p.Name
}

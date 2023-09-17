package main

import "fmt"

type User struct {
	Name string
	Age  string
}

func (u User) Print() {
	fmt.Printf("value is +%v \n", u)
}

func main() {
	user := User{Name: "HNazmul Hassan", Age: "20"}
	user.Print()
}

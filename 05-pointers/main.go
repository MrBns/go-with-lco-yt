package main

import "fmt"

func pointer(value *int) {
	fmt.Printf("address of pointer is %v \n", value)
	fmt.Printf("address of pointer is %v \n", *value)
	*value = 40
}

func main() {
	my_number := 10
	var ptr *int = &my_number
	fmt.Println("Pointer in Golang")
	pointer(ptr)

	fmt.Printf("Edited From another function %v", my_number)
}

package main

import (
	"io/fs"
	"os"
)

func main() {
	path := "./testing.txt"
	file, _ := os.Create(path)

	error := os.WriteFile(path, []byte("My Name is Nazmul Hassan"), fs.ModePerm)
	if error != nil {
		println("Panicked")
		panic(error)
	}

	data := make([]byte, 100)
	file.Read(data)
	println(string(data))

	defer file.Close()
}

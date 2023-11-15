package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["rs"] = "Rust"
	languages["js"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"

	// fmt.Println(languages)
	println("Js stand for %v", languages["js"])
	delete(languages, "rb")

	for key, value := range languages {
		fmt.Printf("%v === %v\n", key, value)
	}
}

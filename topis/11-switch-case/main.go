package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randNumber := rand.Intn(6) + 1

	fmt.Printf("Random Number is %v \n", randNumber)

	switch randNumber {
	case 1:
		fmt.Println("You got 1. Open it now")
	case 2:
		fmt.Println("You got 2.")
	case 3:
		fmt.Println("You got 3.")
	case 4:
		fmt.Println("You got 4.")
	case 5:
		fmt.Println("You got 5.")
	case 6:
		fmt.Println("You got 6.")
		fallthrough
	default:
		fmt.Println("What was that !!!!!!!")
	}
}

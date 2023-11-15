package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

func printTypeAndValue(data any) {
	Printf("Printi data is %v", data)
	Printf("Data is %T", data)
}

func main() {
	Println("Welcome To Our Fizz app")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	inputInNumber, error := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if error != nil {
		Printf(error.Error())
	}
	printTypeAndValue(inputInNumber)

}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func takeInput(sl *[]string) (bool, error) {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return false, err
	} else {

		*sl = append(*sl, input)
		return true, nil
	}
}

func main() {
	var userInputs = []string{}

	for i := 0; i < 9; i++ {
		fmt.Printf("Give your %v Input !\n", i)
		takeInput(&userInputs)
	}

	for index, value := range userInputs {
		fmt.Printf("Value %v is :: %v ", index, value)
	}
}

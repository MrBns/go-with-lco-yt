package userinput

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	println("Welcome To our Shop")
	fmt.Println("Enter Rating for our Pizza")

	//Comma Ok || Err Syntax
	value, _ := reader.ReadString('\n')

	println("thanks for ratting %v", value)
}

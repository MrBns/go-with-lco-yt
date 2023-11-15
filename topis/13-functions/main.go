package main

import "fmt"

func main() {
	value := calculator(sum, 10, 20)
	fmt.Printf("%v \n", value)

	value = restSum(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
	fmt.Printf("%v \n", value)
}
func sum(a, b int) int {
	return a + b
}

func restSum(numbers ...int) (count int) {
	count = 0
	for _, value := range numbers {
		// fmt.Printf("Value is %v \n", value)
		count = sum(count, value)
	}
	return
}

type callback func(int, int) int

func calculator(operation callback, a, b int) int {
	return operation(a, b)
}

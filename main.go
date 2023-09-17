package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	// GetAllTodo()
	MakeNewTodo()
}

func checkNilError(e error) {
	if e != nil {
		print("Panicked")
		panic(e)
	}
}

func extractData[C any](data C, err error) C {
	if err != nil {
		panic(err)
	}
	return data
}

func GetAllTodo() {
	var toDos []Todo
	const url = "https://jsonplaceholder.typicode.com/todos"

	response, err := http.Get(url)
	checkNilError(err)

	// responseBuilder := &strings.Builder{}

	data, err := io.ReadAll(response.Body)
	checkNilError(err)

	print(data)
	// responseBuilder.Write(data)

	json.Unmarshal(data, &toDos)

	for _, todo := range toDos {
		fmt.Printf("%+v \n", todo)
	}
	// println(toDos)

	dataByte, err := json.Marshal(&toDos)
	checkNilError(err)

	fmt.Printf("%v", string(dataByte))
	// fmt.Println(string(dataByte))
	defer response.Body.Close()

}

func MakeNewTodo() {
	todo := &Todo{
		UserId:    1,
		Id:        200,
		Title:     "This is my Custom Todo",
		Completed: false,
	}
	strByte, err := json.Marshal(todo)
	checkNilError(err)

	respond, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", bytes.NewReader(strByte))
	checkNilError(err)

	data := extractData[[]byte](io.ReadAll(respond.Body))

	fmt.Printf("%+v \n", string(data))

	defer respond.Body.Close()
}

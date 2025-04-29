package main

import (
	"fmt"
	"golangDockerDemo/parser"
)

func main() {
	json_data := parser.Json_payload{}
	message, err := json_data.Read_from_file("complex.json")
	json_data.Parse_json(message)

	if err != nil {
		panic(err)
	}

	fmt.Println(message)

}

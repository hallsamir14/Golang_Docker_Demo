package main

import (
	"fmt"
	"golangDockerDemo/parser"
	"os"
)

func main() {

	var filename string = "simple.json"
	var args []string = os.Args

	if len(args) > 1 && args[1] != "" {
		filename = args[1]
	}

	json_data := parser.Json_payload{}
	message, err := json_data.Read_from_file(filename)
	json_data.Parse_json(message)

	if err != nil {
		panic(err)
	}

	fmt.Println(json_data)
}

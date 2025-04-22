package main

import (
	"golangDockerDemo/parser"
)

func main() {
	json_data := parser.Json_payload{}
	message, err := json_data.Get_file("complex.json")
	json_data.Parse_json(message)

	if err != nil {
		panic(err)
	}
}

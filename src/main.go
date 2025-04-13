package main

import (
	"fmt"
	"golangDockerDemo/request_parser"
)

func main() {
	fmt.Println("Lets Parse A Json RQ")
	var message string = `{"name":"Samir"}`
	json_data := request_parser.Json_payload{}
	json_data.Parse_json(message)
}

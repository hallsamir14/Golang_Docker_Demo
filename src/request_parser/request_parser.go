package request_parser

import (
	"encoding/json"
	"fmt"
)

type Json_payload struct {
	Num_of_fields int
	Message_size  int
	Json_data     map[string]interface{}
}

func (j *Json_payload) Parse_json(message string) {
	err := json.Unmarshal([]byte(message), &j.Json_data)

	if err != nil {
		fmt.Println("Error Occured:", err)
	} else {
		fmt.Println("Json Parsed Successfuly")

		j.Num_of_fields = len(j.Json_data)
		j.Message_size = len(message)

		fmt.Println("Parsed Data", j.Json_data)
	}
}

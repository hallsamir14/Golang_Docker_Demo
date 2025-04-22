package struct_parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Json_payload struct {
	num_of_fields int
	message_size  int
	json_data     map[string]interface{}
}

func (j *Json_payload) Parse_json(message string) map[string]interface{} {
	err := json.Unmarshal([]byte(message), &j.json_data)

	if err != nil {
		log.Fatal("Error Occured:", err)
		return nil
	} else {
		log.Print("Json Parsed Successfuly")

		j.num_of_fields = len(j.json_data)
		j.message_size = len(message)

		log.Print("Number of Fields in Payload:", j.num_of_fields)
		log.Print("Byte Size of Payload:", j.message_size)
		log.Print("Parsed Data:", j.json_data)
		return j.json_data
	}
}

func (j *Json_payload) Get_file(filename string) (string, error) {
	file_data, err := os.ReadFile(fmt.Sprintf("../samples/json/%v", filename))

	if err != nil {
		return "", err
	}
	fmt.Println(string(file_data))
	return string(file_data), nil

}

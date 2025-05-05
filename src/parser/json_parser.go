package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
struct will only have 3 data members
num_of_fields and message_size basic metadeta
json_data will store deserialized data from json
*/
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
		log.Print("JSON Parsed Successfuly")

		j.num_of_fields = len(j.json_data)
		j.message_size = len(message)

		log.Print("Number of Fields in Payload:", j.num_of_fields)
		log.Print("Byte Size of Payload:", j.message_size)

		return j.json_data
	}
}

func (j *Json_payload) Read_from_file(filename string) (string, error) {
	//TODO make samples path dynamic so it can read files releative to execution location of binary
	file_data, err := os.ReadFile(fmt.Sprintf("./samples/%v", filename))

	if err != nil {
		return "", err
	}
	log.Print("JSON file read successfully")
	return string(file_data), nil

}

package parser

import (
	"testing"
)

func test_init(t *testing.T) {
	var filename string = "simple.json"

	json_data := Json_payload{}
	message, err := json_data.Read_from_file(filename)
	json_data.Parse_json(message)

	if err != nil {
		t.Error()
	}
}

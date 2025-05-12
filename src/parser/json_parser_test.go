package parser

import (
	"testing"
)

func Test_init(t *testing.T) {
	var message string = `{
        "name": "John Doe",
        "age": 30,
        "email": "johndoe@example.com"
    }`

	json_data := Json_payload{}
	json_data.Parse_json(message)
}

// Package nrp network request protocol in the JSON format
package nrp

import (
	"encoding/json"
	"log"
)

// Simple data type for making simple network requests in JSON format
type Simple struct {
	Id   int         `json:"id,omitempty"`
	Get  string      `json:"get,omitempty"`
	Post string      `json:"post,omitempty"`
	Body interface{} `json:"body,omitempty"`
}

// Export returns a byte slice with the contents of the data structure in JSON format
func (simple *Simple) Export() []byte {
	requestJSON, err := json.Marshal(simple)
	if err != nil {
		log.Println(err)
		return nil
	}
	return requestJSON
}

// Parse fills the structure with the data received in the argument in JSON format
func (simple *Simple) Parse(data []byte) {
	err := json.Unmarshal(data, &simple)
	if err != nil {
		log.Println(err)
	}
}

// GetBody returns the request body as a byte slice in JSON format
func (simple *Simple) GetBody() []byte {
	body, err := json.Marshal(simple.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	return body
}

// BodyToVariable converts data of type map[string]interface{} from body and fills with them the variable referenced in the argument
func (simple *Simple) BodyToVariable(body interface{}) {
	err := json.Unmarshal(simple.GetBody(), body)
	if err != nil {
		log.Println(err)
	}
}

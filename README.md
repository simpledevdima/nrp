# nrp
network request protocol in the JSON format

## Installation
```
go get github.com/skvdmt/nrp
```

## Example
```go
package main

import (
	"fmt"
	"github.com/skvdmt/nrp"
)

type Info struct {
	Valid	bool	`json:"valid"`
	Count	uint8	`json:"count,omitempty"`
}

type Page struct {
	Title	string	`json:"title,omitempty"`
	Content	string	`json:"content,omitempty"`
	Info	Info	`json:"info"`
}

func main() {
	// make page1
	page1 := Page{
		Title: "title",
		Info: Info{
			Valid: false,
			Count: 4,
		},
	}

	// make request
	request1 := nrp.Simple{
		Post: "send",
		Body: page1,
	}
	fmt.Printf("New Struct: %+v\n\n", request1)

	// get request data to export
	requestJSON := request1.Export()
	fmt.Printf("data in JSON: %s\n\n", string(request1.Export()))

	// import request from JSON
	var request2 nrp.Simple
	request2.Parse(requestJSON)
	fmt.Printf("Import Struct: %+v\n\n", request2)

	// get request body
	body := request2.GetBody()
	fmt.Printf("body in JSON: %+v\n\n", string(body))

	// make page 2
	var page2 Page
	request2.BodyToVariable(&page2)
	fmt.Printf("Struct with import data: %+v\n", page2)
}
```
package main

import (
	"encoding/json"

	"github.com/extism/go-pdk"
)

type Query struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

//export ServeDNS
func ServeDNS() int32 {
	input := pdk.Input()

	query := Query{}
	json.Unmarshal(input, &query)

	var output = `{
    "a": [
        {
            "ttl": 300,
            "ip": "5.5.5.5"
        },
        {
            "ttl": 300,
            "ip": "5.5.5.6"
        }
    ]
}`

	pdk.OutputString(output)
	return 0
}

func main() {}

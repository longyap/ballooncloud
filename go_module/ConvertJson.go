package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func main() {
	// XML document to convert to JSON
	xmlDoc := `
        <person>
            <name>John Doe</name>
            <age>35</age>
            <address>
                <street>Main St.</street>
                <city>Anytown</city>
                <country>USA</country>
            </address>
        </person>
    `

	// Parse XML into a map
	var personMap map[string]interface{}
	err := xml.Unmarshal([]byte(xmlDoc), &personMap)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	// Marshal map to JSON
	jsonData, err := json.Marshal(personMap)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	// Print JSON data
	fmt.Println(string(jsonData))
}

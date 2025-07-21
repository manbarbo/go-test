package utils

import (
	"encoding/json"
	"fmt"
)

// PrintDataAsJSON: prints struct o slice de structs as formated JSON
func PrintDataAsJSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error convirtiendo a JSON:", err)
		return
	}
	fmt.Println("JSON formateado:")
	fmt.Println(string(jsonData))
}

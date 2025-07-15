package utils

import (
	"encoding/json"
	"fmt"
)

// PrintDataAsJSON imprime cualquier struct o slice de structs como JSON formateado
func PrintDataAsJSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error convirtiendo a JSON:", err)
		return
	}
	fmt.Println("JSON formateado:")
	fmt.Println(string(jsonData))
}

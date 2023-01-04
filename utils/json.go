package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJSON(filename string) map[string][]string {
	var parsedData map[string][]string

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error to access file!")
		os.Exit(1)
	}
	json.Unmarshal(data, &parsedData)

	return parsedData
}

func WriteJSON(filename string, data map[string][]string) error {
	parsedData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	os.WriteFile(filename, parsedData, 0755)
	return nil
}

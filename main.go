/*
Copyright © 2022 allandlobr allandlobr@gmail.com
*/
package main

import (
	"os"

	"github.com/allandlobr/go-bookmark-cli/cmd"
)

func readJSON(filename string) error {
	_, err := os.ReadFile(filename)

	if os.IsNotExist(err) {
		os.WriteFile("db.json", []byte("{\"main\": \"[]\"}"), 0755)
	}

	return nil
}

func main() {
	os.Setenv("DB_FILENAME", "db.json")
	err := readJSON("db.json")
	if err != nil {
		return
	}
	cmd.Execute()
}

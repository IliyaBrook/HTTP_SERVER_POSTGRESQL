package data

import (
	"HTTP_SERVER/utils"
	"encoding/json"
	"fmt"
	"os"
)

var filePath = utils.ResolvePath("data/database.json")

func (d *DB) ReadDatabase() error {

	file, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	if err := json.Unmarshal(file, d); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return err
	}

	return nil
}

func (d *DB) SaveDatabase() error {
	file, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filePath, file, 0644); err != nil {
		return err
	}

	return nil
}

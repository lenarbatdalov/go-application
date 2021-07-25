package commands

import (
	"log"
	"os"
)

func DefineDatabase() {
	fileName := "database/database.sqlite"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
}

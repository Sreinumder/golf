package touch

import (
	"log"
	"os"
)

func Touch(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}

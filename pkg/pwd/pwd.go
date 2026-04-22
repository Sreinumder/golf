package pwd

import (
	"fmt"
	"os"
)

func Pwd() (string, error) {
	if pwdir, err := os.Getwd(); err == nil {
		return pwdir, nil
	} else {
		fmt.Println(err)
		return "", err
	}
	// return os.Getwd()
}

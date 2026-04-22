package main

import (
	"fmt"

	"rojgosai.com.np/golf/pkg/pwd"
)

func main() {
	if dir, err := pwd.Pwd(); err == nil {
		fmt.Print("Current pwd:", dir)
	} else {
		print(err)
	}

}

package cat

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Cat(filename string, flags map[string]int) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if flag := flags["n"]; flag > 0 {
		content, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		lines := "1\t" + string(content)
		lineArr := []rune(lines) // used for representing characters
		result := ""
		lineNum := 2
		for _, ch := range lineArr {
			result += string(ch)
			if ch == '\n' {
				result += fmt.Sprintf("%d", lineNum) + "\t"
				lineNum++
			}
		}
		// Handle case where file does not end with newline
		if len(lineArr) > 0 && lineArr[len(lineArr)-1] != '\n' {
			result += "{" + fmt.Sprintf("%d", lineNum) + "}"
		}
		os.Stdout.WriteString(result)
	} else {
		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			log.Fatal(err)
		}
	}
}

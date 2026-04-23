package main

import (
	"fmt"
	"os"
	"strings"

	"rojgosai.com.np/golf/pkg/cat"
	"rojgosai.com.np/golf/pkg/ls"
	"rojgosai.com.np/golf/pkg/pwd"
)

func parseFlags(args []string) (map[string]int, []string) {
	flags := make(map[string]int)
	var paths []string
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") && len(arg) > 2 {
			flag := arg[2:]
			flags[flag]++
		} else if strings.HasPrefix(arg, "-") && len(arg) > 1 {
			for _, ch := range arg[1:] {
				flags[string(ch)]++
			}
		} else {
			paths = append(paths, arg)
		}
	}
	return flags, paths
}

func main() {
	args := os.Args

	switch args[1] {

	case "pwd":
		fmt.Println(pwd.Pwd())

	case "ls":
		// Parse flags and paths
		flags, paths := parseFlags(args[2:])
		if len(paths) > 0 {
			for _, path := range paths {
				fmt.Println("\n" + path + ":")
				ls.Ls(path, flags)
			}
		} else {
			fmt.Println(".:")
			ls.Ls(".", flags)
		}

	case "cat":
		flags, paths := parseFlags(args[2:])
		if len(paths) > 0 {
			for _, path := range paths {
				fmt.Println("\n" + path + ":")
				cat.Cat(path, flags)
			}
		} else {
			fmt.Println("error: atleast one argument required.")
		}
	}
}

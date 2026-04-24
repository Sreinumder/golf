package main

import (
	"fmt"
	"os"
	"strings"

	"rojgosai.com.np/golf/lf"
	"rojgosai.com.np/golf/pkg/cat"
	"rojgosai.com.np/golf/pkg/ls"
	"rojgosai.com.np/golf/pkg/mv"
	"rojgosai.com.np/golf/pkg/pwd"
	"rojgosai.com.np/golf/pkg/touch"
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
	case "touch":
		_, paths := parseFlags(args[2:])
		for _, paths := range paths {
			touch.Touch(paths)
		}
	case "mv":
		_, paths := parseFlags(args[2:])
		mv.Mv(paths)
	case "lf":
		_, paths := parseFlags(args[2:])
		if len(paths) > 0 {
			lf.Lf(paths[0])
		} else {
			lf.Lf(".")
		}
	}

}

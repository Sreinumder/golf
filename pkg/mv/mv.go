package mv

import (
	"fmt"
	"os"
	"strings"
)

func Mv(path_args []string) {
	args_count := len(path_args)

	switch args_count {
	case 0:
		fmt.Println("mv: missing operand")
		os.Exit(1)
	case 1:
		fmt.Printf("mv: missing destination argument after %s\n", path_args[0])
		os.Exit(1)
	case 2:
		if strings.HasSuffix(path_args[1], "/") {
			if info, err := os.Stat(path_args[1]); err != nil {
				if os.IsNotExist(err) {
					fmt.Printf("the %s destination directory does not exist", path_args[1])
				} else {
					fmt.Printf("some error on ", path_args[1])
					os.Exit(1)
				}
			} else if info.IsDir() {
				err := os.Rename(path_args[0], path_args[1]+path_args[0])
				if err != nil {
					fmt.Print("there was some errors: while moving ", err)
					os.Exit(1)
				}
			}
		} else {
			err := os.Rename(path_args[0], path_args[1])
			if err != nil {
				fmt.Print("there was some errors: while renaming ", err)
				os.Exit(1)
			}
		}
	default:
		length := len(path_args)
		last_arg := path_args[length-1]
		if strings.HasSuffix(last_arg, "/") {
			if info, err := os.Stat(last_arg); err != nil {
				if os.IsNotExist(err) {
					fmt.Printf("the %s destination directory does not exist", last_arg)
				} else {
					fmt.Printf("some error on ", last_arg)
				}
			} else if info.IsDir() {
				// implementing this is kinda hard ngl
				// for i, arg
				// err := os.Rename(path_args[0], last_arg+path_args[0])
				// if err != nil {
				// 	fmt.Print("there was some errors: while moving ", err)
				// }
			}
		} else {
			fmt.Print("the last argument needs to be an destination")
			os.Exit(1)
		}
	}
}

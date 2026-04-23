package ls

import (
	"fmt"
	"os"
)

func Ls(path string, flags map[string]int) {
	// Define allowed flags (single-letter and word flags)
	allowed := []string{"l", "long"}
	var passed_allowed_flags []string
	for flag, _ := range flags {
		notPresent := true
		for _, allowedFlag := range allowed {
			if flag == allowedFlag {
				passed_allowed_flags = append(passed_allowed_flags, flag)
				notPresent = false
			}
		}
		if notPresent {
			fmt.Println("The flag ", flag, "does not exist for this cmd")
			os.Exit(2)
		}
	}

	// Validate flags
	for flag := range flags {
		switch flag {
		case "l", "long":
			// allowed
		default:
			fmt.Fprintf(os.Stderr, "Unknown flag: %s\n", flag)
			os.Exit(1)
		}
	}

	if entries, err := os.ReadDir(path); err == nil {
		for _, entry := range entries {
			name := entry.Name()
			if entry.IsDir() {
				name += "/"
			}
			if entryInfo, EntryErr := entry.Info(); EntryErr == nil {
				if boo := flags["l"]; boo > 0 {
					fmt.Println(entryInfo.Mode(), "\t", entryInfo.Size(), "\t", entryInfo.ModTime(), "\t", name)
				} else {
					fmt.Print(name, "\t")
				}
			} else {
				fmt.Print(EntryErr)
			}
		}
	} else {
		fmt.Println("Error occured: ", err)
	}
}

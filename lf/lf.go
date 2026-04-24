package lf

import (
	"os"

	"github.com/rivo/tview"
)

func Lf(path string) {
	app := tview.NewApplication()
	if entries, err := os.ReadDir(path); err == nil {
		// box := tview.NewBox().SetBorder(true).SetTitle(path)
		list := tview.NewList()
		for i, entry := range entries {
			list.AddItem(entry.Name(), entry.Name(), rune(i), func() { print("hello world") })
		}
		if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
			panic(err)
		}
		// if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		// 	panic(err)
		// }
	} else {
		print("there was some err to get pwd: ", err)
		os.Exit(1)
	}
}

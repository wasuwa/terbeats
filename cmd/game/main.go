package main

import (
	"github.com/rivo/tview"
	ui "github.com/wasuwa/terbeats/internal/ui"
)

func main() {
	app := tview.NewApplication()
	root := ui.NewGame(app)
	if err := app.SetRoot(root, true).Run(); err != nil {
		panic(err)
	}
}

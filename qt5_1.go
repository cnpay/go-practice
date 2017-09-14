package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("hello world")
	window.SetMinimumSize2(400, 300)
	window.Show()
	app.Exec()
}

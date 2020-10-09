package main

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

const (
	KEY_LEFT  uint = 65361
	KEY_UP    uint = 65362
	KEY_RIGHT uint = 65363
	KEY_DOWN  uint = 65364
)

func main() {
	gtk.Init(nil)

	// gui boilerplate
	window, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	drawArea, _ := gtk.DrawingAreaNew()
	window.Add(drawArea)
	window.SetTitle("小方块")
	window.Connect("destroy", gtk.MainQuit)
	window.SetDefaultSize(800, 600)
	window.ShowAll()

	// Data
	unitSize := 20.0
	x := 0.0
	y := 0.0
	keyMap := map[uint]func(){
		KEY_LEFT:  func() { x-- },
		KEY_UP:    func() { y-- },
		KEY_RIGHT: func() { x++ },
		KEY_DOWN:  func() { y++ },
	}

	// Event handlers
	drawArea.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		cr.SetSourceRGB(0, 0, 0)
		cr.Rectangle(x*unitSize, y*unitSize, unitSize, unitSize)
		cr.Fill()
	})

	window.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{ev}
		if move, found := keyMap[keyEvent.KeyVal()]; found {
			move()
			win.QueueDraw()
		}
	})

	gtk.Main()
}

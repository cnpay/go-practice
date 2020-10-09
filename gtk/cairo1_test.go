package main

import (
	"testing"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func Test_cario1(t *testing.T) {
	const (
		_KeyLeft  uint = 65361
		_KeyUp    uint = 65362
		_KeyRight uint = 65363
		_KeyDown  uint = 65364
	)
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
		_KeyLeft:  func() { x-- },
		_KeyUp:    func() { y-- },
		_KeyRight: func() { x++ },
		_KeyDown:  func() { y++ },
	}

	// Event handlers
	drawArea.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		cr.SetSourceRGB(0, 0, 0)
		cr.Rectangle(x*unitSize, y*unitSize, unitSize, unitSize)
		cr.Fill()
	})

	window.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		keyEvent := &gdk.EventKey{Event: ev}
		if move, found := keyMap[keyEvent.KeyVal()]; found {
			move()
			win.QueueDraw()
		}
	})

	gtk.Main()
}

package test

import (
	"testing"

	"github.com/webview/webview"
)

func Test_webview1(t *testing.T) {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("一个webview的测试")
	w.SetSize(768, 1024, webview.HintNone)
	w.Navigate("https://m.baidu.com")
	w.Run()
}

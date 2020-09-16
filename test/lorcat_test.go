package test

import (
	"log"
	"testing"

	"github.com/zserge/lorca"
)

func Test_lorcat1(t *testing.T) {
	/* 	ui, err := lorca.New("data:text/html,"+url.PathEscape(`
	   	<html>
	   		<head><title>Hello</title></head>
	   		<body><h1>Hello, world!</h1></body>
	   	</html>
	   	`), "", 480, 320) */

	ui, err := lorca.New("https://m.baidu.com", "", 480, 600)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}

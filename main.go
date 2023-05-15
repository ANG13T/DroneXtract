package main

// OSINT, DAT File parsing, other file format parsing, crystal web server, TUI and GUI 
// "github.com/TwiN/go-color"	

import (
	"fmt"
	"io/ioutil"
	"github.com/iskaa02/qalam/gradient"
)

func Banner() {
	banner, _ := ioutil.ReadFile("txt/banner.txt")
	info, _ := ioutil.ReadFile("txt/info.txt")
	// options, _ := ioutil.ReadFile("txt/options.txt")
	g,_:=gradient.NewGradient("cyan", "blue")
	solid,_:=gradient.NewGradient("blue", "#1179ef")
	// opt,_:=gradient.NewGradient("#1179ef", "cyan")
	g.Print(string(banner))
	solid.Print(string(info))
	// opt.Print("\n" + string(options))
}

func main() {
	fmt.Println("\n")
	Banner()
}
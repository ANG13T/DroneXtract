package main

// OSINT, DAT File parsing, other file format parsing, crystal web server, TUI and GUI

import (
	"fmt"
	"io/ioutil"
	"github.com/iskaa02/qalam/gradient"
	"github.com/TwiN/go-color"
	"github.com/ANG13T/SatIntel/osint"
	"strconv"
	"os"
)

func Banner() {
	banner, _ := ioutil.ReadFile("txt/banner.txt")
	info, _ := ioutil.ReadFile("txt/info.txt")
	options, _ := ioutil.ReadFile("txt/options.txt")
	g,_:=gradient.NewGradient("cyan", "blue")
	solid,_:=gradient.NewGradient("blue", "#1179ef")
	opt,_:=gradient.NewGradient("#1179ef", "cyan")
	g.Print(string(banner))
	solid.Print(string(info))
	opt.Print("\n" + string(options))
}

func main() {
	fmt.Println("hello world")
	Banner()
}
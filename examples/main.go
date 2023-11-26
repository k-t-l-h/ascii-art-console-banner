package main

import (
	"banner"
	"github.com/fatih/color"
)
func main() {

	renderer := banner.NewAsciiRender( banner.WithColor(color.FgCyan), banner.WithBoldness(true))
	renderer.Render("ktlh")


}


package main

import (
	"os"

	"dev.lopo.mail-support/ui"
)

func init() {
	os.Setenv("FYNE_THEME", "light") //makes the ui in ligth mode
	os.Setenv("FYNE_SCALE", "2")     //makes the ui in ligth mode
}

func main() {
	ui.RenderUI()
}

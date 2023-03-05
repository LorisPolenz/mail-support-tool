package main

import (
	"os"

	"dev.lopo.oma/helper"
	"dev.lopo.oma/ui"
	// "dev.lopo.oma/ui"
)

func init() {
	os.Setenv("FYNE_THEME", "light") //makes the ui in ligth mode
	os.Setenv("FYNE_SCALE", "2")     //makes the ui in ligth mode
}

func main() {
	helper.ValidateEnVariables()
	ui.RenderUI()
}

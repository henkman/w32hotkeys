package main

import (
	w32hk "github.com/papplampe/w32hotkeys"
)

var (
	ctrl1id int
	ctrl2id int
)

func ControlNumPressed(id int) {
	println("one of the control + number keys was pressed")
	switch id {
	case ctrl1id:
		println("ctrl + 1 was pressed")
	case ctrl2id:
		println("ctrl + 2 was pressed")
	}
}

func AltEPressed(id int) {
	println("alt + E was pressed")
}

func main() {
	hotkeys := w32hk.New()
	hotkeys.AddHotkey(w32hk.ALT|w32hk.NOREPEAT, 'E', AltEPressed)
	ctrl1id = hotkeys.AddHotkey(w32hk.CONTROL|w32hk.NOREPEAT, '1', ControlNumPressed)
	ctrl2id = hotkeys.AddHotkey(w32hk.CONTROL|w32hk.NOREPEAT, '2', ControlNumPressed)
	hotkeys.Start()
	println("hotkeys now active")
	select {}
}

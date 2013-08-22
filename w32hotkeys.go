package w32hotkey

// #include <windows.h>
import "C"

import (
	"runtime"
)

const (
	ALT      uint = 0x0001
	CONTROL  uint = 0x0002
	SHIFT    uint = 0x0004
	WIN      uint = 0x0008
	NOREPEAT uint = 0x4000
)

type Hotkey struct {
	ModKey   C.UINT
	Vk       C.UINT
	Callback Callback
}

type Hotkeys struct {
	hotkeys map[C.int]*Hotkey
	id      C.int
}

type Callback func(id int)

func New() *Hotkeys {
	hotkeys := new(Hotkeys)
	hotkeys.hotkeys = make(map[C.int]*Hotkey)
	hotkeys.id = 0
	return hotkeys
}

func (this *Hotkeys) AddHotkey(ModKey, Vk uint, callback Callback) int {
	hotkey := new(Hotkey)
	hotkey.ModKey = C.UINT(ModKey)
	hotkey.Vk = C.UINT(Vk)
	hotkey.Callback = callback
	this.hotkeys[this.id] = hotkey
	defer func() {
		this.id++
	}()
	return int(this.id)
}

func (this *Hotkeys) Start() {
	go this.run()
}

func (this *Hotkeys) run() {
	runtime.LockOSThread()

	for id, hk := range this.hotkeys {
		if C.RegisterHotKey(nil, id, hk.ModKey, hk.Vk) == 0 {
			panic("could not register hotkey")
		}
	}

	for {
		var msg C.MSG
		C.GetMessage((*C.struct_tagMSG)(&msg), nil, 0, 0)
		if msg.message == C.WM_HOTKEY {
			id := C.int(msg.wParam)
			hk, ok := this.hotkeys[id]
			if ok {
				go hk.Callback(int(id))
			}
		}
	}
}
